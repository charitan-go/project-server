package project

import (
	"github.com/charitan-go/project-server/internal/project/handler"
	"github.com/charitan-go/project-server/internal/project/repository"
	"github.com/charitan-go/project-server/internal/project/service"
	"go.uber.org/fx"
)

var ProjectModule = fx.Module("project",
	fx.Provide(
		handler.NewProjectHandler,
		service.NewProjectService,
		repository.NewProjectRepository,
	),
)
