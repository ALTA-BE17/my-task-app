package router

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	dep "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type Routes struct {
	dig.In
	Dep     dep.Dependency
	User    user.UserHandler // interface handler
	Project project.ProjectHandler
	// Task    task.TaskHandler
}

func (r *Routes) RegisterRoutes() {
	e := r.Dep.Echo
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.POST("/register", r.User.Register())
	e.POST("/login", r.User.Login())
	e.GET("/users", r.User.Profile(), jwt.JWT([]byte(config.JWTSECRET)))
	e.GET("/users/search", r.User.SearchUsers(), jwt.JWT([]byte(config.JWTSECRET)))
	e.PUT("/users", r.User.UpdateProfile(), jwt.JWT([]byte(config.JWTSECRET)))
	e.DELETE("/users", r.User.Deactive(), jwt.JWT([]byte(config.JWTSECRET)))

	e.POST("/projects", r.Project.CreateProject(), jwt.JWT([]byte(config.JWTSECRET)))
	// e.GET("/projects", r.Project.ListProjects(), jwt.JWT([]byte(config.JWTSECRET))) // include dengan get list task by project id dan get username
	// e.GET("/projects/:id", r.Project.GetProjectByProjectID(), jwt.JWT([]byte(config.JWTSECRET)))
	// e.PUT("/projects/:id", r.Project.UpdateProjectByProjectID(), jwt.JWT([]byte(config.JWTSECRET)))
	// e.DELETE("/projects/:id", r.Project.DeleteProjectByProjectID(), jwt.JWT([]byte(config.JWTSECRET)))

	// e.POST("/tasks/:id", r.Task.CreateTaskByTaskID(), jwt.JWT([]byte(config.JWTSECRET)))
	// e.PUT("/tasks/:id", r.Task.UpdateTaskByTaskID(), jwt.JWT([]byte(config.JWTSECRET)))
	// e.DELETE("/tasks/:id", r.Task.DeleteTaskByTaskID(), jwt.JWT([]byte(config.JWTSECRET)))
}
