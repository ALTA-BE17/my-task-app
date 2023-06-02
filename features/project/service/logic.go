package service

import (
	"errors"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	query    project.ProjectData
	validate *validator.Validate
}

func New(pd project.ProjectData) project.ProjectService {
	return &Service{
		query:    pd,
		validate: validator.New(),
	}
}

// CreateProject implements project.ProjectService
func (s *Service) CreateProject(userId string, request project.Core) (project.Core, error) {
	result, err := s.query.CreateProject(userId, request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "duplicated") {
			message = "data already used, duplicated"
		} else {
			message = "internal server error"
		}
		return project.Core{}, errors.New(message)
	}

	return result, nil
}

// DeleteProjectByProjectID implements project.ProjectService
func (*Service) DeleteProjectByProjectID(userId string, projectId string) (project.Core, error) {
	panic("unimplemented")
}

// GetProjectByProjectID implements project.ProjectService
func (*Service) GetProjectByProjectID(userId string, projectId string) (project.Core, error) {
	panic("unimplemented")
}

// ListProjects implements project.ProjectService
func (*Service) ListProjects(userId string) ([]project.Core, error) {
	panic("unimplemented")
}

// UpdateProjectByProjectID implements project.ProjectService
func (*Service) UpdateProjectByProjectID(userId string, projectId string, request project.Core) (project.Core, error) {
	panic("unimplemented")
}
