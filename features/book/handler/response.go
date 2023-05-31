package handler

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
)

type ListBookResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Author    string       `json:"author"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      ResponseUser `json:"user"`
}

type ResponseUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetBookResponse(book book.Core) ListBookResponse {
	response := ListBookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
		User: ResponseUser{
			Name:  book.User.Name,
			Email: book.User.Email,
		},
	}

	return response
}
