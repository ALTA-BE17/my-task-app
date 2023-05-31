package handler

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
)

type SearchUsersResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SearchUsers(repo user.Core) SearchUsersResponse {
	return SearchUsersResponse{
		ID:        repo.ID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Name:      repo.Name,
		Email:     repo.Email,
	}
}

type ProfileResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Profile(repo user.Core) ProfileResponse {
	return ProfileResponse{
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Name:      repo.Name,
		Email:     repo.Email,
	}
}

type GetAllUsersResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Books     []ResponseBook `json:"books"`
}

type ResponseBook struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllUsers(user user.Core) GetAllUsersResponse {
	response := GetAllUsersResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Books:     make([]ResponseBook, len(user.Books)),
	}

	for i, book := range user.Books {
		response.Books[i] = ResponseBook{
			Title:     book.Title,
			Author:    book.Author,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}
	}

	return response
}
