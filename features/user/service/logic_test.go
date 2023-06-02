package service

import (
	"errors"
	"testing"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper/validation"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProfile(t *testing.T) {
	data := mocks.NewUserData(t)
	usr := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@gmail.com", Password: "HashedPassword1"}
	request := user.Core{Username: "newName", Phone: "081235288543", Email: "newemail@gmail.com", Password: "HashedPassword1", NewPassword: "HashedPassword2", ConfirmedPassword: "HashedPassword2"}
	service := New(data)

	t.Run("name, email, and password cannot be empty", func(t *testing.T) {
		arguments := user.Core{Password: "", NewPassword: "newPassword", ConfirmedPassword: "newPassword"}
		_, err := service.UpdateProfile("550e8400-e29b-41d4-a716-446655440000", arguments)
		expectedErr := errors.New("it's not a complete request for updating the password")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("old password and current password do not match", func(t *testing.T) {
		request := user.Core{
			Password:          "HashedPassword1",
			NewPassword:       "HashedPassword2",
			ConfirmedPassword: "HashedPassword2",
		}

		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(usr, nil)
		match1 := helper.MatchPassword(request.Password, usr.Password)
		assert.False(t, match1, "Expected passwords to not match")

		_, err := service.UpdateProfile("550e8400-e29b-41d4-a716-446655440000", request)
		expectedErr := errors.New("old password and current password do not match")
		assert.NotNil(t, err, "Expected an error")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("new password and confirmed password do not match", func(t *testing.T) {
		request := user.Core{
			Password:          "HashedPassword1",
			NewPassword:       "HashedPassword",
			ConfirmedPassword: "HashedPassword2",
		}

		match2 := helper.MatchNewPassword(request.NewPassword, request.ConfirmedPassword)
		expectedErr := errors.New("new password and confirmed password do not match")
		assert.False(t, match2, "Expected passwords to not match")
		assert.EqualError(t, expectedErr, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("password strength is low", func(t *testing.T) {
		request := user.Core{
			Password:          "HashedPassword1",
			NewPassword:       "Hashed",
			ConfirmedPassword: "Hashed",
		}
		err := validation.UpdatePasswordValidator(request.NewPassword)
		expectedErr := errors.New("password strength is low")
		assert.NotNil(t, err)
		assert.Error(t, err, "Expected validation error")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("empty input", func(t *testing.T) {
		request := user.Core{}
		_, err := service.UpdateProfile("550e8400-e29b-41d4-a716-446655440000", request)
		expectedErr := errors.New("failed to process the request due to empty input")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("error while retrieving user profile", func(t *testing.T) {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(user.Core{}, errors.New("error while retrieving user profile"))
		_, err := service.UpdateProfile("550e8400-e29b-41d4-a716-446655440000", request)
		expectedErr := errors.New("error while retrieving user profile")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("failed to update user, duplicate data entry", func(t *testing.T) {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(usr, nil)
		data.On("UpdateProfile", "550e8400-e29b-41d4-a716-446655440000", usr).Return(user.Core{}, errors.New("failed to update user, duplicate data entry"))
		_, err := service.UpdateProfile("550e8400-e29b-41d4-a716-446655440000", usr)
		expectedErr := errors.New("failed to update user, duplicate data entry")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("internal server error, failed to update account", func(t *testing.T) {
		expectedError := errors.New("internal server error, failed to update account")
		userId := "550e8400-e29b-41d4-a716-446655440000"
		data.On("Profile", userId).Return(user.Core{}, nil).Once()
		data.On("UpdateProfile", userId, usr).Return(user.Core{}, expectedError).Once()

		result, err := service.UpdateProfile(userId, usr)

		assert.Equal(t, user.Core{}, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error, failed to update account")

		data.AssertExpectations(t)
	})
}

func BenchmarkSearchUsers(b *testing.B) {
	data := mocks.NewUserData(b)
	pattern := "admin"
	result := []user.Core{{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@mail.com"}}
	service := New(data)

	data.On("SearchUsers", mock.AnythingOfType("string"), pattern).Return(result, nil).Times(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.SearchUsers("550e8400-e29b-41d4-a716-446655440000", pattern)
		assert.Nil(b, err)
	}

	data.AssertExpectations(b)
}

func TestSearchUsers(t *testing.T) {
	data := mocks.NewUserData(t)
	pattern := "admin"
	emptyPattern := ""
	result := []user.Core{{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@mail.com"}}
	service := New(data)

	t.Run("failed to process the request due to empty param input", func(t *testing.T) {
		_, err := service.SearchUsers("550e8400-e29b-41d4-a716-446655440000", emptyPattern)
		expectedErr := errors.New("failed to process the request due to empty param input")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "failed to process the request due to empty param input")
		data.AssertExpectations(t)
	})

	t.Run("success search users account", func(t *testing.T) {
		data.On("SearchUsers", "550e8400-e29b-41d4-a716-446655440000", pattern).Return(result, nil).Once()
		res, err := service.SearchUsers("550e8400-e29b-41d4-a716-446655440000", pattern)
		assert.Nil(t, err)
		assert.Equal(t, len(result), len(res))
		data.AssertExpectations(t)
	})

	t.Run("error while searching list user", func(t *testing.T) {
		data.On("SearchUsers", "550e8400-e29b-41d4-a716-446655440000", pattern).Return([]user.Core{}, errors.New("not found, error while searching list user")).Once()
		res, err := service.SearchUsers("550e8400-e29b-41d4-a716-446655440000", pattern)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found, error while retrieving list users")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("SearchUsers", "550e8400-e29b-41d4-a716-446655440000", pattern).Return([]user.Core{}, errors.New("internal server error")).Once()
		res, err := service.SearchUsers("550e8400-e29b-41d4-a716-446655440000", pattern)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, []user.Core{}, res)
		data.AssertExpectations(t)
	})
}

func BenchmarkDeactive(b *testing.B) {
	data := mocks.NewUserData(b)
	result := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@mail.com"}
	service := New(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data.On("Deactive", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		_, _ = service.Deactive("550e8400-e29b-41d4-a716-446655440000")
	}
	data.AssertExpectations(b)
}

func TestDeactive(t *testing.T) {
	data := mocks.NewUserData(t)
	result := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@mail.com"}
	service := New(data)

	t.Run("success deactivate an account", func(t *testing.T) {
		data.On("Deactive", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		_, err := service.Deactive("550e8400-e29b-41d4-a716-446655440000")
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("account not registered", func(t *testing.T) {
		data.On("Deactive", "550e8400-e29b-41d4-a716-446655440000").Return(user.Core{}, errors.New("error while retrieving user profile")).Once()
		res, err := service.Deactive("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "account not registered")
		assert.Equal(t, user.Core{}, res)
		data.AssertExpectations(t)
	})

	t.Run("internal server error, failed to delete account", func(t *testing.T) {
		data.On("Deactive", "550e8400-e29b-41d4-a716-446655440000").Return(user.Core{}, errors.New("internal server error, failed to delete account")).Once()
		res, err := service.Deactive("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, user.Core{}, res)
		data.AssertExpectations(t)
	})
}

func BenchmarkProfile(b *testing.B) {
	data := mocks.NewUserData(b)
	result := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@mail.com"}
	service := New(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		_, err := service.Profile("550e8400-e29b-41d4-a716-446655440000")
		if err != nil {
			b.Errorf("Unexpected error: %s", err)
		}
		data.AssertExpectations(b)
	}
}

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)
	result := user.Core{
		UserID:   "550e8400-e29b-41d4-a716-446655440000",
		Username: "admin",
		Phone:    "081235288543",
		Email:    "admin@mail.com"}
	service := New(data)

	t.Run("success get profile", func(t *testing.T) {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(result, nil).Once()
		res, err := service.Profile("550e8400-e29b-41d4-a716-446655440000")
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.Equal(t, result.Username, res.Username)
		assert.Equal(t, result.Phone, res.Phone)
		assert.Equal(t, result.Email, res.Email)
		data.AssertExpectations(t)
	})

	t.Run("not found, error while retrieving user profile", func(t *testing.T) {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(user.Core{}, errors.New("not found, error while retrieving user profile")).Once()
		res, err := service.Profile("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, user.Core{}, res)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Profile", "550e8400-e29b-41d4-a716-446655440000").Return(user.Core{}, errors.New("internal server error")).Once()
		res, err := service.Profile("550e8400-e29b-41d4-a716-446655440000")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Equal(t, user.Core{}, res)
		data.AssertExpectations(t)
	})
}

func BenchmarkLogin(b *testing.B) {
	data := mocks.NewUserData(b)
	arguments := user.Core{Username: "grace", Password: "@SecretPassword123"}
	hashed, _ := helper.HashPassword(arguments.Password)
	resultData := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "grace", Password: hashed}
	service := New(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data.On("Login", mock.Anything).Return(resultData, "", nil).Once()
		_, _, err := service.Login(arguments)
		if err != nil {
			b.Errorf("Unexpected error: %s", err)
		}
		data.AssertExpectations(b)
	}
}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	arguments := user.Core{Username: "admin", Password: "@SecretPassword123"}
	wrongArguments := user.Core{Username: "admin", Password: "@WrongPassword"}
	token := "123"
	emptyToken := ""
	hashed, _ := helper.HashPassword(arguments.Password)
	result := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Password: hashed}
	service := New(data)

	t.Run("username and password cannot be empty", func(t *testing.T) {
		request := user.Core{
			Username: "",
			Password: "",
		}

		_, _, err := service.Login(request)
		expectedErr := errors.New("username and password cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success login", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(result, token, nil).Once()
		res, token, err := service.Login(arguments)
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, result.Username, res.Username)
		data.AssertExpectations(t)
	})

	t.Run("account has not been registered", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.Core{}, emptyToken, errors.New("account not registered")).Once()
		_, _, err := service.Login(wrongArguments)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "account has not been registered")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Login", mock.Anything).Return(user.Core{}, emptyToken, errors.New("server error")).Once()
		res, token, err := service.Login(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.UserID)
		assert.Equal(t, emptyToken, token)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func BenchmarkRegister(b *testing.B) {
	data := mocks.NewUserData(b)
	arguments := user.Core{
		Username: "admin",
		Phone:    "081235288543",
		Email:    "admin@gmail.com",
		Password: "@S3#cr3tP4ss#word123",
	}
	resultData := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@gmail.com"}
	service := New(data)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data.On("Register", mock.Anything).Return(resultData, nil).Once()
		_, err := service.Register(arguments)
		if err != nil {
			b.Errorf("Unexpected error: %s", err)
		}
		data.AssertExpectations(b)
	}
}

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	arguments := user.Core{
		Username: "admin",
		Phone:    "081235288543",
		Email:    "admin@gmail.com",
		Password: "@S3#cr3tP4ss#word123",
	}
	result := user.Core{UserID: "550e8400-e29b-41d4-a716-446655440000", Username: "admin", Phone: "081235288543", Email: "admin@gmail.com"}
	service := New(data)

	t.Run("request cannot be empty", func(t *testing.T) {
		request := user.Core{
			Username: "admin",
			Phone:    "081235288543",
			Email:    "",
			Password: "",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("request cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("field validation for password", func(t *testing.T) {
		request := user.Core{
			Username: "admin",
			Phone:    "081235288543",
			Email:    "admin@gmail.com",
			Password: "",
		}

		_, err := validation.UserValidate("register", request)
		assert.Error(t, err, "Expected validation error")

		expectedErr := errors.New("Key: 'Register.Password' Error:Field validation for 'Password' failed on the 'required' tag")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("invalid email format", func(t *testing.T) {
		request := user.Core{
			Username: "admin",
			Phone:    "081235288543",
			Email:    "admin.mail.com",
			Password: "@S3#cr3tP4ss#word123",
		}

		_, err := service.Register(request)

		expectedError := errors.New("invalid email format")
		assert.EqualError(t, err, expectedError.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("check password strength, low password", func(t *testing.T) {
		request := user.Core{
			Username: "admin",
			Phone:    "081235288543",
			Email:    "admin@gmail.com",
			Password: "123",
		}

		_, err := service.Register(request)

		expectedError := errors.New("check password strength, low password")
		assert.EqualError(t, err, expectedError.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success create account", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(result, nil).Once()
		res, err := service.Register(arguments)
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.NotEmpty(t, result.Username)
		assert.NotEmpty(t, result.Phone)
		assert.NotEmpty(t, result.Email)
		data.AssertExpectations(t)
	})

	t.Run("data already used, duplicated", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("data already used, duplicated")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.UserID)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.UserID)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}
