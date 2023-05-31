package book

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `validate:"required"`
	Author    string `validate:"required"`
	UserID    uint
	User      User
}

type User struct {
	Name  string
	Email string
}

type BookHandler interface {
	InsertBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	ListBooks() echo.HandlerFunc
}

type BookService interface {
	InsertBook(userId uint, request Core) (Core, error)
	UpdateBook(userId uint, request *Core) (Core, error)
	ListBooks() ([]Core, error)
}

type BookData interface {
	InsertBook(userId uint, request Core) (Core, error)
	UpdateBook(userId uint, request *Core) (Core, error)
	ListBooks() ([]Core, error)
}
