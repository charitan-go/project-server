package service

import (
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
}

type projectServiceImpl struct {
	r repository.ProjectRepository
}

func NewProjectService(r repository.ProjectRepository) ProjectService {
	return &projectServiceImpl{r}
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
	// Validation for date

	// Create project dto to save to db
	saveProjectEntDto := req.ToSaveProjectEntDto(jwtPayload)

	projectDto, err := svc.r.Save(saveProjectEntDto)
	if err != nil {
		log.Printf("Cannot save project to db: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusInternalServerError, Message: "Cannot create project."}
	}

	responseDto := svc.toProjectResponseDto(projectDto)

	return responseDto, nil
}

func (svc *projectServiceImpl) HandleGetProjectByIdRest(
	projectIdStr string,
) (*dto.ProjectResponseDto, *dto.ErrorResponseDto) {
	log.Printf("Project id: %s", projectIdStr)
	projectId, err := uuid.Parse(projectIdStr)
	if err != nil {
		log.Printf("Cannot convert str to id: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Project not found."}
	}

	// Find project by id
	projectDto, err := svc.r.FindOneByReadableId(projectId)
	if err != nil {
		log.Printf("Cannot find project: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusBadRequest, Message: "Project not found."}
	}

	responseDto := svc.toProjectResponseDto(projectDto)

	return responseDto, nil
}
