package data

import (
	"errors"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
	"go.uber.org/zap"
)

type Query struct {
	dep dependency.Dependency
}

func New(dep dependency.Dependency) project.ProjectData {
	return &Query{dep: dep}
}

// CreateProject implements project.ProjectData
func (q *Query) CreateProject(userId string, request project.Core) (project.Core, error) {
	createdProject := project.Core{}
	req := projectEntities(request)
	req.UserID = userId
	trx := q.dep.DB.Begin()
	if trx.Error != nil {
		q.dep.Logger.Warn("failed to begin database transaction")
		return project.Core{}, errors.New("internal server error")
	}

	if err := trx.Table("projects").Create(&req).Error; err != nil {
		q.dep.Logger.Warn("rollback while create project.")
		trx.Rollback()
		if strings.Contains(err.Error(), "Error 1452 (23000)") {
			q.dep.Logger.Warn("failed to insert data, user does not exist", zap.Error(err))
			return project.Core{}, errors.New("user does not exist")
		}
		if strings.Contains(err.Error(), "Error 1062 (23000)") {
			q.dep.Logger.Warn("failed to insert data, duplicate entry", zap.Error(err))
			return project.Core{}, errors.New("duplicate entry")
		}

		q.dep.Logger.Warn("failed to insert data, error while creating transaction", zap.Error(err))
		return project.Core{}, errors.New("internal server error")
	}

	commitErr := trx.Commit().Error
	if commitErr != nil {
		q.dep.Logger.Warn("rollback while commit transaction.")
		trx.Rollback()
		q.dep.Logger.Warn("failed to commit transaction", zap.Error(commitErr))
		return project.Core{}, errors.New("internal server error, failed to commit database transaction")
	}

	q.dep.Logger.Sugar().Infof("new book has been created: %s, %s", createdProject.Name)
	return createdProject, nil
}

// DeleteProjectByProjectID implements project.ProjectData
func (*Query) DeleteProjectByProjectID(userId string, projectId string) (project.Core, error) {
	panic("unimplemented")
}

// GetProjectByProjectID implements project.ProjectData
func (*Query) GetProjectByProjectID(userId string, projectId string) (project.Core, error) {
	panic("unimplemented")
}

// ListProjects implements project.ProjectData
func (*Query) ListProjects(userId string) ([]project.Core, error) {
	panic("unimplemented")
}

// UpdateProjectByProjectID implements project.ProjectData
func (*Query) UpdateProjectByProjectID(userId string, projectId string, request project.Core) (project.Core, error) {
	panic("unimplemented")
}
