package api

import (
	"net/http"

	projecthandler "github.com/charitan-go/project-server/internal/project/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	ProjectHandler *projecthandler.ProjectHandler
}

func NewApi(projectHandler *projecthandler.ProjectHandler) *Api {
	return &Api{
		ProjectHandler: projectHandler,
	}
}

func (*Api) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
