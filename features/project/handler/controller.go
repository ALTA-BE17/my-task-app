package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type ProjectHandler struct {
	dig.In
	Service project.ProjectService
}

func New(srv project.ProjectService) project.ProjectHandler {
	return &ProjectHandler{
		Service: srv,
	}
}

// CreateProject implements project.ProjectService
func (ph *ProjectHandler) CreateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := CreateProjectRequest{}

		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, err1 := ph.Service.CreateProject(userId, *RequestToCore(request))
		if err1 != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Username, email, and password cannot be empty.", nil))
			}
			if strings.Contains(err.Error(), "duplidate") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Invalid email format", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error, duplicate data entry.", nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created a project.", nil))
	}
}
