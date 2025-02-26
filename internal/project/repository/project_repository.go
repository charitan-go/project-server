package repository

import (
	"github.com/charitan-go/project-server/internal/project/model"
	"github.com/charitan-go/project-server/pkg/database"
	"gorm.io/gorm"
	"log"
)

type ProjectRepository interface {
	Save(projectModel *model.Project) (*model.Project, error)

	FindOneByEmail(email string) (*model.Project, error)
	FindOneByReadableId(readableId string) (*model.Project, error)
}

type projectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository() ProjectRepository {
	db := database.DB
	if db == nil {
		log.Println("db is nil")
	} else {
		log.Println("db is not nil")
	}

	return &projectRepositoryImpl{db: database.DB}
}

func (r *projectRepositoryImpl) Save(projectModel *model.Project) (*model.Project, error) {
	result := r.db.Create(projectModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return projectModel, nil
}

func (r *projectRepositoryImpl) FindOneByEmail(email string) (*model.Project, error) {
	var project model.Project

	result := r.db.Where("email = ?", email).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

// FindOneByReadableId implements ProjectRepository.
func (r *projectRepositoryImpl) FindOneByReadableId(readableId string) (*model.Project, error) {
	var project model.Project

	result := r.db.Where("readable_id = ?", readableId).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}
