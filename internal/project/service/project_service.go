package service

import (
	"log"
	"net/http"

	"github.com/charitan-go/project-server/ent"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/internal/project/repository"
)

type ProjectService interface {
	HandleCreateProjectRest(req *dto.CreateProjectRequestDto) (*dto.ProjectResponseDto, *dto.ErrorResponseDto)
}

type projectServiceImpl struct {
	r repository.ProjectRepository
}

func NewProjectService(r repository.ProjectRepository) ProjectService {
	return &projectServiceImpl{r}
}

func (svc *projectServiceImpl) toProjectResponseDto(p *ent.Project) *dto.ProjectResponseDto {
	return &dto.ProjectResponseDto{
		ReadableId:  p.ReadableID.String(),
		Name:        p.Name,
		Description: p.Description,
		Goal:        p.Goal,
		StartDate:   p.StartDate.UnixMilli(),
		EndDate:     p.EndDate.UnixMilli(),
		Category:    string(p.Category),
		CountryCode: p.CountryCode,
	}
}

// HandleCreateProjectRest implements ProjectService.
func (svc *projectServiceImpl) HandleCreateProjectRest(req *dto.CreateProjectRequestDto) (*dto.ProjectResponseDto, *dto.ErrorResponseDto) {
	// TODO: Impl
	saveProjectEntDto := req.ToSaveProjectEntDto()

	projectDto, err := svc.r.Save(saveProjectEntDto)
	if err != nil {
		log.Fatalf("Cannot save project to db: %v\n", err)
		return nil, &dto.ErrorResponseDto{StatusCode: http.StatusInternalServerError, Message: "Cannot create project."}
	}

	responseDto := svc.toProjectResponseDto(projectDto)

	return responseDto, nil
}
