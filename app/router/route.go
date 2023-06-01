package router

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	dep "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type Routes struct {
	dig.In
	Dep  dep.Dependency
	User user.UserHandler // interface handler
	Book book.BookHandler
}

func (r *Routes) RegisterRoutes() {
	e := r.Dep.Echo
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError: true,
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogHost:     true,
		LogLatency:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			r.Dep.Logger.Info("request",
				zap.String("host", v.Host),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("remote_ip", v.RemoteIP),
				zap.Duration("latency", v.Latency),
			)
			return nil
		},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.POST("/register", r.User.Register())
	e.POST("/login", r.User.Login())
	e.GET("/users/books", r.User.GetAllUsers())
	e.GET("/users", r.User.Profile(), jwt.JWT([]byte(config.JWTSecret)))
	e.GET("/users/search", r.User.SearchUsers(), jwt.JWT([]byte(config.JWTSecret)))
	e.PUT("/users", r.User.UpdateProfile(), jwt.JWT([]byte(config.JWTSecret)))
	e.DELETE("/users", r.User.Deactive(), jwt.JWT([]byte(config.JWTSecret)))

	e.GET("/books", r.Book.ListBooks())
	e.POST("/books", r.Book.InsertBook(), jwt.JWT([]byte(config.JWTSecret)))
	e.PUT("/books/:id", r.Book.UpdateBook(), jwt.JWT([]byte(config.JWTSecret)))
}
