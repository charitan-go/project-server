package service

import ()

type ProjectService interface {
}

type projectServiceImpl struct {
}

func NewProjectService() ProjectService {
	return &projectServiceImpl{}
}
