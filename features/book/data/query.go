package data

import (
	"errors"
	"fmt"
	"log"

	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	"go.uber.org/zap"
)

type Query struct {
	dep dependency.Dependency
}

func New(dep dependency.Dependency) book.BookData {
	return &Query{dep: dep}
}

// CreateBook implements book.Data
func (q *Query) InsertBook(userId uint, request book.Core) (book.Core, error) {
	createdBook := book.Core{}
	req := bookEntities(request)
	req.UserID = userId
	var countRowAffected int64
	transaction := q.dep.DB.Begin()
	if transaction.Error != nil {
		q.dep.Logger.Warn("failed to begin database transaction")
		return book.Core{}, errors.New("internal server error, failed to begin database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	if err := transaction.Table("books").Create(&req).Count(&countRowAffected).Error; err != nil {
		q.dep.Logger.Warn("rollback while create book.")
		transaction.Rollback()
		if strings.Contains(err.Error(), "Error 1452 (23000)") {
			q.dep.Logger.Warn("failed to insert data, user does not exist", zap.Error(err))
			return book.Core{}, errors.New("user does not exist")
		}
		if strings.Contains(err.Error(), "Error 1062 (23000)") {
			q.dep.Logger.Warn("failed to insert data, duplicate entry", zap.Error(err))
			return book.Core{}, errors.New("duplicate entry")
		}

		q.dep.Logger.Warn("failed to insert data, error while creating transaction", zap.Error(err))
		return book.Core{}, errors.New("internal server error, failed to create database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	commitErr := transaction.Commit().Error
	if commitErr != nil {
		q.dep.Logger.Warn("rollback while commit transaction book.") //gunakan fatal untuk os.exit(1)
		transaction.Rollback()
		q.dep.Logger.Warn("failed to commit transaction", zap.Error(commitErr))
		return book.Core{}, errors.New("internal server error, failed to commit database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	q.dep.Logger.Sugar().Infof("new book has been created: %s, %s", createdBook.Title, createdBook.Author)
	return createdBook, nil
}

// UpdateBook implements book.Data
func (q *Query) UpdateBook(userId uint, request *book.Core) (book.Core, error) {
	createdBook := book.Core{}
	req := bookEntities(*request)
	req.UserID = userId
	var countRowAffected int64
	transaction := q.dep.DB.Begin()
	if transaction.Error != nil {
		q.dep.Logger.Warn("failed to begin database transaction")
		return book.Core{}, errors.New("internal server error, failed to begin database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	if err := transaction.Table("books").Where("id = ? AND user_id = ?", req.ID, req.UserID).Updates(req).Count(&countRowAffected).Error; err != nil {
		q.dep.Logger.Warn("rollback while updating book.")
		transaction.Rollback()

		if strings.Contains(err.Error(), "Error 1452 (23000)") {
			q.dep.Logger.Warn("failed to update data, user does not exist", zap.Error(err))
			return book.Core{}, errors.New("user does not exist")
		}

		if strings.Contains(err.Error(), "Error 1062 (23000)") {
			q.dep.Logger.Warn("failed to update data, duplicate entry", zap.Error(err))
			return book.Core{}, errors.New("duplicate entry")
		}

		q.dep.Logger.Warn("failed to update data, error while updating transaction", zap.Error(err))
		return book.Core{}, errors.New("internal server error, failed to update database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	if err := transaction.Commit().Error; err != nil {
		q.dep.Logger.Warn("rollback while committing transaction for updating book.")
		transaction.Rollback()
		q.dep.Logger.Warn("failed to commit transaction", zap.Error(err))
		return book.Core{}, errors.New("internal server error, failed to commit database transaction")
	}

	log.Printf("row affected : %d\n", countRowAffected)

	q.dep.Logger.Sugar().Infof("book has been updated: %s, %s", createdBook.Title, createdBook.Author)
	return createdBook, nil
}

// ListBooks implements book.BookData
func (q *Query) ListBooks() ([]book.Core, error) {
	var books []Book

	// gunakan preload
	query := q.dep.DB.Table("books").Preload("User").Find(&books)
	if query.Error != nil {
		q.dep.Logger.Sugar().Infof("failed to list books: %v", query.Error)
		return nil, query.Error
	}

	var bookCores []book.Core
	for _, b := range books {
		bookCores = append(bookCores, bookModels(b))
	}

	fmt.Println(bookCores)
	return bookCores, nil
}
