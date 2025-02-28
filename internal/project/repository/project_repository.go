package repository

import (
	"context"
	"log"

	"github.com/charitan-go/project-server/ent"
	"github.com/charitan-go/project-server/ent/project"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/pkg/database"
	"github.com/google/uuid"
)

type ProjectRepository interface {
	// Save(projectModel *model.Project) (*model.Project, error)
	Save(*dto.SaveProjectEntDto) (*ent.Project, error)

	FindAll(page, size int) ([]*ent.Project, error)
	FindOneByReadableId(readableId uuid.UUID) (*ent.Project, error)
}

type projectRepositoryImpl struct {
	client *ent.Client
	// db *gorm.DB
}

func NewProjectRepository() ProjectRepository {
	// db := database.DB
	client := database.Client
	if client == nil {
		log.Println("db is nil")
	} else {
		log.Println("db is not nil")
	}

	return &projectRepositoryImpl{client}
}

func (r *projectRepositoryImpl) Save(entDto *dto.SaveProjectEntDto) (*ent.Project, error) {
	project, err := r.client.Project.
		Create().
		SetName(entDto.Name).
		SetDescription(entDto.Description).
		SetGoal(entDto.Goal).
		SetStartDate(entDto.StartDate).
		SetEndDate(entDto.EndDate).
		SetCategory(entDto.Category).
		SetCountryCode(entDto.CountryCode).
		SetOwnerCharityReadableID(entDto.OwnerCharityReadableId).
		SetStatus(entDto.Status).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return project, nil
}

// FindAll implements ProjectRepository.
func (r *projectRepositoryImpl) FindAll(page int, size int) ([]*ent.Project, error) {
	projectList, err := r.client.Project.
		Query().
		Offset((page - 1) * size).
		Order(
			project.ByStartDate(),
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return projectList, err
}

func (r *projectRepositoryImpl) FindOneByReadableId(readableId uuid.UUID) (*ent.Project, error) {
	project, err := r.client.Project.
		Query().
		Where(project.ReadableID(readableId)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return project, nil
}

// func (r *projectRepositoryImpl) FindOneByEmail(email string) (*model.Project, error) {
// 	var project model.Project
//
// 	result := r.db.Where("email = ?", email).First(&project)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
//
// 	return &project, nil
// }
//
// // FindOneByReadableId implements ProjectRepository.
// func (r *projectRepositoryImpl) FindOneByReadableId(readableId string) (*model.Project, error) {
// 	var project model.Project
//
// 	result := r.db.Where("readable_id = ?", readableId).First(&project)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
//
// 	return &project, nil
// }
