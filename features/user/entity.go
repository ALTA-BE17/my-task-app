package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	UserID            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
	Username          string
	Phone             string
	Email             string
	Password          string
	NewPassword       string
	ConfirmedPassword string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Login() echo.HandlerFunc
	SearchUsers() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	Deactive() echo.HandlerFunc
}

type UserService interface {
	Register(request Core) (Core, error)
	Login(request Core) (Core, string, error)
	Profile(userId string) (Core, error)
	SearchUsers(userId string, pattern string) ([]Core, error)
	UpdateProfile(userId string, request Core) (Core, error)
	Deactive(userId string) (Core, error)
}

type UserData interface {
	Register(request Core) (Core, error)
	Login(request Core) (Core, string, error)
	Profile(userId string) (Core, error)
	SearchUsers(userId string, pattern string) ([]Core, error)
	UpdateProfile(userId string, request Core) (Core, error)
	Deactive(userId string) (Core, error)
}
