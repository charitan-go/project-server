package repository

import (
	"context"
	"log"

	"github.com/charitan-go/project-server/ent"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/pkg/database"
)

type ProjectRepository interface {
	// Save(projectModel *model.Project) (*model.Project, error)
	Save(*dto.SaveProjectEntDto) (*ent.Project, error)
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
		SetCategory(dto.CategoryEnum(entDto.Category)).
		SetCountryCode(entDto.CountryCode).
		Save(context.Background())

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
