package handler

import (
	"log"
	"net/http"

	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/charitan-go/project-server/internal/project/service"
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
	req := new(dto.CreateProjectRequestDto)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request body", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleCreateProjectRest(req)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

// func (h *ProjectHandler) RegisterDonor(c echo.Context) error {
// 	req := new(dto.RegisterDonorRequestDto)
// 	if err := c.Bind(req); err != nil {
// 		log.Println(err)
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request bodyy", StatusCode: http.StatusBadRequest})
// 	}
//
// 	res, errRes := h.svc.HandleRegisterDonorRest(req)
// 	if errRes != nil {
// 		return c.JSON(int(errRes.StatusCode), *errRes)
// 	}
//
// 	return c.JSON(http.StatusCreated, *res)
// }
