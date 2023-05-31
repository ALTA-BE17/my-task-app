package data

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey; autoIncrement"`
	Title  string `gorm:"type:VARCHAR(100);not null;unique"`
	Author string `gorm:"type:VARCHAR(100);not null"`
	UserID uint
	User   User
}

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; autoIncrement"`
	Name     string `gorm:"type:VARCHAR(100);not null;unique"`
	Email    string `gorm:"type:VARCHAR(100);not null;unique"`
	Password string `gorm:"type:VARCHAR(225);not null"`
	Books    []Book `gorm:"foreignKey:UserID"`
}

// note: solusi model user diletakkan di model book, efek redudansi.
// user/book dihandle dari sisi query book , select * from book.where user id = ?
//. yang dimiliki user1,

// Book-model to book-core
func bookModels(repo Book) book.Core {
	return book.Core{
		ID:        repo.ID,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Title:     repo.Title,
		Author:    repo.Author,
		UserID:    repo.UserID,
	}
}

// Book-core to book-model
func bookEntities(repo book.Core) Book {
	return Book{
		Model: gorm.Model{
			ID:        repo.ID,
			CreatedAt: repo.CreatedAt,
			UpdatedAt: repo.UpdatedAt,
		},
		Title:  repo.Title,
		Author: repo.Author,
	}
}
