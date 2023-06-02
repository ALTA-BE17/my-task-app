package project

import (
	"time"

	task "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/task"
	"github.com/labstack/echo/v4"
)

type Core struct {
	ProjectID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	UserID      string
	Username    string
	Tasks       []task.Core
}

type ProjectHandler interface {
	CreateProject() echo.HandlerFunc
	// ListProjects() echo.HandlerFunc
	// GetProjectByProjectID() echo.HandlerFunc
	// UpdateProjectByProjectID() echo.HandlerFunc
	// DeleteProjectByProjectID() echo.HandlerFunc
}

type ProjectService interface {
	CreateProject(userId string, request Core) (Core, error)
	ListProjects(userId string) ([]Core, error)
	GetProjectByProjectID(userId string, projectId string) (Core, error)
	UpdateProjectByProjectID(userId string, projectId string, request Core) (Core, error)
	DeleteProjectByProjectID(userId string, projectId string) (Core, error)
}

type ProjectData interface {
	CreateProject(userId string, request Core) (Core, error)
	ListProjects(userId string) ([]Core, error)
	GetProjectByProjectID(userId string, projectId string) (Core, error)
	UpdateProjectByProjectID(userId string, projectId string, request Core) (Core, error)
	DeleteProjectByProjectID(userId string, projectId string) (Core, error)
}
