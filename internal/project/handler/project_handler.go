package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/internal/project/service"
	restpkg "github.com/charitan-go/project-server/pkg/rest"
	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	svc service.ProjectService
}

func (h *ProjectHandler) CheckHealth() string {
	return "OK"
}

func NewProjectHandler(svc service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

func (h *ProjectHandler) CreateProject(c echo.Context) error {
	jwtPayload, err := restpkg.GetJwtPayload(c)
	if err != nil {
		log.Println("Not found header payload")
		return c.JSON(http.StatusNonAuthoritativeInfo, dto.ErrorResponseDto{Message: "Not authorized"})
	}

	req := new(dto.CreateProjectRequestDto)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request body", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleCreateProjectRest(req, jwtPayload)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *ProjectHandler) GetProjectById(c echo.Context) error {
	projectId := c.Param("projectId")

	res, errRes := h.svc.HandleGetProjectByIdRest(projectId)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	pageStr, sizeStr := c.Param("page"), c.Param("size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		if pageStr != "" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
		}
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		if sizeStr != "" {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
		}
		size = 10
	}

	if page <= 0 || size <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request param", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleGetProjectsRest(page, size)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, res)
}
