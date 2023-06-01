package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	UserID            uint
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
	Profile(userId uint) (Core, error)
	SearchUsers(userId uint, pattern string) ([]Core, error)
	UpdateProfile(userId uint, request Core) (Core, error)
	Deactive(userId uint) (Core, error)
}

type UserData interface {
	Register(request Core) (Core, error)
	Login(request Core) (Core, string, error)
	Profile(userId uint) (Core, error)
	SearchUsers(userId uint, pattern string) ([]Core, error)
	UpdateProfile(userId uint, request Core) (Core, error)
	Deactive(userId uint) (Core, error)
}
