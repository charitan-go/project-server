package handler

import (
	"github.com/charitan-go/project-server/internal/project/service"
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
