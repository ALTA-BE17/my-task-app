package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type UserHandler struct {
	dig.In
	Service user.UserService
	Dep     dependency.Dependency
}

func New(srv user.UserService, dep dependency.Dependency) user.UserHandler {
	return &UserHandler{
		Service: srv,
		Dep:     dep,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, err := uh.Service.Register(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Username, email, and password cannot be empty.", nil))
			}
			if strings.Contains(err.Error(), "low password") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Low strength password, at least 60%", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error, duplicate data entry.", nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created an account.", nil))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := LoginRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, token, err := uh.Service.Login(RequestToCore(request))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Authentication failed due to incorrect username or password.", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successful login", map[string]interface{}{"token": token}))
	}
}

func (uh *UserHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		user, err := uh.Service.Profile(userId)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving user profile.", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error.", nil))
		}

		profileResp := Profile(user)
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", profileResp))
	}
}

func (uh *UserHandler) SearchUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		pattern := c.QueryParam("pattern")
		if pattern == "" {
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Failed to process the request due to empty input.", nil))
		}
		users, err := uh.Service.SearchUsers(userId, pattern)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving list user.", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error.", nil))
			}
		}

		result := make([]SearchUsersResponse, len(users))
		for i, user := range users {
			result[i] = SearchUsers(user)
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", result))
	}
}

func (uh *UserHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		request := UpdateProfileRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to process the request.", nil))
		}

		_, err = uh.Service.UpdateProfile(userId, RequestToCore(&request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request: data cannot be empty while updating, at least one field must be provided for update", nil))
			}
			if strings.Contains(err.Error(), "it's not a complete request for updating the password") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request: it's not a complete request for updating the password", nil))
			}
			if strings.Contains(err.Error(), "old password and current password do not match") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request: old password and current password do not match", nil))
			}
			if strings.Contains(err.Error(), "new password and confirmed password do not match") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request: new password and confirmed password do not match", nil))
			}
			if strings.Contains(err.Error(), "password strength is low") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request: Low strength password, at least 60%", nil))
			}
			if strings.Contains(err.Error(), "failed to update user") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusBadRequest, "Error due to duplicated entry, data has been used by other user.", nil))
			}
			if strings.Contains(err.Error(), "error while retrieving user profile") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving user profile.", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Successfully updated an account", nil))
	}
}

func (uh *UserHandler) Deactive() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil))
		}

		_, err1 := uh.Service.Deactive(userId)
		if err1 != nil {
			if strings.Contains(err.Error(), "error while retrieving user profile") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving user profile.", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Successfully deleted an account", nil))
	}
}

// GetAllUsers implements user.UserHandler
func (uh *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.Service.GetAllUserHasBooks()
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "Error while retrieving list user.", nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error.", nil))
			}
		}

		result := make([]GetAllUsersResponse, len(users))
		for i, user := range users {
			result[i] = GetAllUsers(user)
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", result))
	}
}
