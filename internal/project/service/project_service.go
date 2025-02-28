package service

import (
	"context"
	"log"
	"net/http"

	"github.com/charitan-go/project-server/ent"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/internal/project/repository"
	restpkg "github.com/charitan-go/project-server/pkg/rest"
	"github.com/google/uuid"
)

type ProjectService interface {
	HandleCreateProjectRest(req *dto.CreateProjectRequestDto, jwtPayload *restpkg.JwtPayload) (*dto.ProjectResponseDto, *dto.ErrorResponseDto)

	HandleGetProjectByIdRest(projectId string) (*dto.ProjectResponseDto, *dto.ErrorResponseDto)
	HandleGetProjectsRest(page, size int) ([]*dto.ProjectResponseDto, *dto.ErrorResponseDto)
}

type projectServiceImpl struct {
	r        repository.ProjectRepository
	redisSvc ProjectRedisService
}

func NewProjectService(
	r repository.ProjectRepository,
	redisSvc ProjectRedisService,
) ProjectService {
	return &projectServiceImpl{r, redisSvc}
}

func (svc *projectServiceImpl) toProjectResponseDto(p *ent.Project) *dto.ProjectResponseDto {
	return &dto.ProjectResponseDto{
		ReadableId:             p.ReadableID.String(),
		Name:                   p.Name,
		Description:            p.Description,
		Goal:                   p.Goal,
		StartDate:              p.StartDate.UnixMilli(),
		EndDate:                p.EndDate.UnixMilli(),
		Category:               string(p.Category),
		CountryCode:            p.CountryCode,
		Status:                 string(p.Status),
		OwnerCharityReadableId: p.OwnerCharityReadableID,
	}
}

// HandleCreateProjectRest implements ProjectService.
func (svc *projectServiceImpl) HandleCreateProjectRest(
	req *dto.CreateProjectRequestDto,
	jwtPayload *restpkg.JwtPayload,
) (*dto.ProjectResponseDto, *dto.ErrorResponseDto) {
	ctx := context.Background()

	// Validation for date

	// Create project dto to save to db
	saveProjectEntDto := req.ToSaveProjectEntDto(jwtPayload)

	projectDto, err := svc.r.Save(saveProjectEntDto)
	if err != nil {
		log.Printf("Cannot save project to db: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusInternalServerError, Message: "Cannot create project."}
	}

	// Cache to redis
	err = svc.redisSvc.SetById(ctx, projectDto)
	if err != nil {
		log.Printf("Cannot save to redis: %v\n", err)
	}

	responseDto := svc.toProjectResponseDto(projectDto)

	return responseDto, nil
}

func (svc *projectServiceImpl) HandleGetProjectByIdRest(
	projectIdStr string,
) (*dto.ProjectResponseDto, *dto.ErrorResponseDto) {
	ctx := context.Background()

	log.Printf("Project id: %s", projectIdStr)
	projectId, err := uuid.Parse(projectIdStr)
	if err != nil {
		log.Printf("Cannot convert str to id: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Project not found."}
	}

	// TODO: Find in redis first
	p, err := svc.redisSvc.GetById(ctx, projectIdStr)
	if err != nil {
		log.Printf("Cannot get by id from redis: %v\n", err)
	}

	// Can find string from redis
	if p != nil {
		log.Println("CACHED HIT!!!")
		responseDto := svc.toProjectResponseDto(p)
		return responseDto, nil
	}
	log.Println("CACHED MISS!!!")

	// Find project by id
	projectDto, err := svc.r.FindOneByReadableId(projectId)
	if err != nil {
		log.Printf("Cannot find project: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Project not found."}
	}

	err = svc.redisSvc.SetById(ctx, projectDto)
	if err != nil {
		log.Printf("Cannot save to redis: %v\n", err)
	}

	responseDto := svc.toProjectResponseDto(projectDto)

	return responseDto, nil
}

// HandleGetProjectsRest implements ProjectService.
func (svc *projectServiceImpl) HandleGetProjectsRest(page int, size int) ([]*dto.ProjectResponseDto, *dto.ErrorResponseDto) {
	// Query project
	projectDtoList, err := svc.r.FindAll(page, size)
	if err != nil {
		log.Printf("Cannot query: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Project not found."}
	}

	// Map to response dto
	responseDto := make([]*dto.ProjectResponseDto, len(projectDtoList))
	for i, dto := range projectDtoList {
		responseDto[i] = svc.toProjectResponseDto(dto)
	}

	return responseDto, nil
}
