package router

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	dep "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type Routes struct {
	dig.In
	Dep  dep.Dependency
	User user.UserHandler // interface handler
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
}
