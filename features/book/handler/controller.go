package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type BookHandler struct {
	dig.In
	Service book.BookService
	Dep     dependency.Dependency
}

func New(srv book.BookService, dep dependency.Dependency) book.BookHandler {
	return &BookHandler{
		Service: srv,
		Dep:     dep,
	}
}

// CreateBook implements book.Handler
func (bh *BookHandler) InsertBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err1 := middlewares.ExtractToken(c)
		log.Println(userId)
		log.Printf("%T", userId)
		if err1 != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		request := InsertBookRequest{}

		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, err := bh.Service.InsertBook(userId, *RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Title and author cannot be empty.", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error, duplicate data entry.", nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created a book.", nil))
	}
}

// UpdateBook implements book.Handler
func (bh *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		bookId, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Failed to process the request due to empty input.", nil))
		}

		request := UpdateBookRequest{ID: uint(bookId)}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, err = bh.Service.UpdateBook(userId, RequestToCore(&request))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))

		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Successfully updated an account", nil))
	}
}

// ListBooks implements book.BookHandler
func (bh *BookHandler) ListBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.Service.ListBooks()
		if err != nil {
			if strings.Contains(err.Error(), "books not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving list books.", nil))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error.", nil))
			}
		}

		result := make([]ListBookResponse, len(books))
		for i, book := range books {
			result[i] = GetBookResponse(book)
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", result))
	}
}
