package service

import (
	"errors"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Service struct {
	query    book.BookData
	validate *validator.Validate
	dep      dependency.Dependency
}

func New(bd book.BookData, dep dependency.Dependency) book.BookService {
	return &Service{
		query:    bd,
		validate: validator.New(),
		dep:      dep,
	}
}

// CreateBook implements book.Service
func (s *Service) InsertBook(userId uint, request book.Core) (book.Core, error) {
	if request.Title == "" || request.Author == "" {
		s.dep.Logger.Info("title and author cannot be empty")
		return book.Core{}, errors.New("title and author cannot be empty")
	}

	err := s.validate.Struct(request)
	if err != nil {

		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field()+" is invalid")
		}

		if strings.Contains(err.Error(), "input should be text") {
			s.dep.Logger.Info("Error while validating new book, check input should be text.", zap.Error(err))
		}

		s.dep.Logger.Info("Error while validating new book, check input should be string.", zap.Error(err))
		return book.Core{}, errors.New(strings.Join(validationErrors, ", "))
	}

	result, err := s.query.InsertBook(userId, request)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			s.dep.Logger.Info("Error while validating new book, check input should be string.", zap.Error(err))
		}
		return book.Core{}, errors.New("internal server error")
	}

	return result, nil
}

// UpdateBook implements book.Service
func (s *Service) UpdateBook(userId uint, request *book.Core) (book.Core, error) {
	if request.Title == "" && request.Author == "" {
		s.dep.Logger.Info("data cannot be empty")
		return book.Core{}, errors.New("failed to process the request due to empty input")
	}

	result, err := s.query.UpdateBook(userId, request)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return book.Core{}, errors.New("failed to update user, duplicate data entry")
		}
		return book.Core{}, errors.New("internal server error")
	}

	return result, nil
}

// ListBooks implements book.BookService
func (s *Service) ListBooks() ([]book.Core, error) {
	result, err := s.query.ListBooks()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return []book.Core{}, errors.New("not found, error while retrieving list books")
		} else {
			return []book.Core{}, errors.New("internal server error")
		}
	}
	return result, nil
}
