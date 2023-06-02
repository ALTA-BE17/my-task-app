package task

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	TaskID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Status      string
	ProjectID   string
}

type TaskHandler interface {
	CreateTaskByTaskID() echo.HandlerFunc
	UpdateTaskByTaskID() echo.HandlerFunc
	DeleteTaskByTaskID() echo.HandlerFunc
}

type TaskService interface {
	CreateTaskByTaskID(userId string, request Core) (Core, error)
	UpdateTaskByTaskID(userId string) ([]Core, error)
	DeleteTaskByTaskID(userId string, projectId string) (Core, error)
}

type TaskData interface {
	CreateTaskByTaskID(userId string, request Core) (Core, error)
	UpdateTaskByTaskID(userId string) ([]Core, error)
	DeleteTaskByTaskID(userId string, projectId string) (Core, error)
}
