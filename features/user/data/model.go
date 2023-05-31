package data

import (
	bookCore "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	book "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/data"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint        `gorm:"primaryKey; autoIncrement"`
	Name     string      `gorm:"type:VARCHAR(100);not null;unique"`
	Email    string      `gorm:"type:VARCHAR(100);not null;unique"`
	Password string      `gorm:"type:VARCHAR(225);not null"`
	Books    []book.Book `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

// User-model to user-core
func userModels(repo User) user.Core {
	return user.Core{
		ID:        repo.ID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Name:      repo.Name,
		Email:     repo.Email,
		Password:  repo.Password,
		Books:     make([]bookCore.Core, 0),
	}
}

// User-core to user-model
func userEntities(repo user.Core) User {
	return User{
		Model: gorm.Model{
			ID:        repo.ID,
			CreatedAt: repo.CreatedAt,
			UpdatedAt: repo.UpdatedAt,
		},
		Name:     repo.Name,
		Email:    repo.Email,
		Password: repo.Password,
	}
}

// Book-model to book-core
func bookModels(repo book.Book) bookCore.Core {
	return bookCore.Core{
		ID:        repo.ID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Title:     repo.Title,
		Author:    repo.Author,
	}
}
