package service

import (
	"errors"
	"log"
	"strings"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper/validation"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	query    user.UserData
	validate *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &Service{
		query:    ud,
		validate: validator.New(),
	}
}

func (s *Service) Register(request user.Core) (user.Core, error) {
	if request.Name == "" || request.Email == "" || request.Password == "" {
		return user.Core{}, errors.New("username, email, and password cannot be empty")
	}

	err := validation.Authenticate(validation.UserValidate("register", request))
	if err != nil {
		return user.Core{}, errors.New("check password strength, low password")
	}

	result, err := s.query.Register(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "duplicated") {
			message = "data already used, duplicated"
		} else {
			message = "internal server error"
		}
		return user.Core{}, errors.New(message)
	}

	return result, nil
}

func (s *Service) Login(request user.Core) (user.Core, string, error) {
	if request.Name == "" || request.Password == "" {
		return user.Core{}, "", errors.New("username and password cannot be empty")
	}

	result, token, err := s.query.Login(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "account not registered") {
			message = "account has not been registered"
		} else {
			message = "internal server error"
		}
		return user.Core{}, "", errors.New(message)
	}

	return result, token, nil
}

func (s *Service) Profile(userId uint) (user.Core, error) {
	result, err := s.query.Profile(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found, error while retrieving user profile") {
			return user.Core{}, errors.New("not found, error while retrieving user profile")
		} else {
			return user.Core{}, errors.New("internal server error")
		}
	}
	return result, nil
}

func (s *Service) SearchUsers(userId uint, pattern string) ([]user.Core, error) {
	if pattern == "" {
		return []user.Core{}, errors.New("failed to process the request due to empty param input")
	}

	result, err := s.query.SearchUsers(userId, pattern)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return []user.Core{}, errors.New("not found, error while retrieving list users")
		} else {
			return []user.Core{}, errors.New("internal server error")
		}
	}
	return result, nil
}

func (s *Service) UpdateProfile(userId uint, request user.Core) (user.Core, error) {
	if request.Password != "" || request.NewPassword != "" || request.ConfirmedPassword != "" {
		if request.Password == "" || request.NewPassword == "" || request.ConfirmedPassword == "" {
			return user.Core{}, errors.New("it's not a complete request for updating the password")
		}

		usr, err := s.query.Profile(userId)
		log.Println(usr)
		if err != nil {
			return user.Core{}, err
		}

		match1 := helper.MatchPassword(request.Password, usr.Password)
		log.Printf("match: %v", match1)
		if !match1 {
			return user.Core{}, errors.New("old password and current password do not match")
		}

		match2 := helper.MatchNewPassword(request.NewPassword, request.ConfirmedPassword)
		if !match2 {
			return user.Core{}, errors.New("new password and confirmed password do not match")
		}

		err = validation.UpdatePasswordValidator(request.NewPassword)
		if err != nil {
			return user.Core{}, errors.New("password strength is low")
		}

		hash, err := helper.HashPassword(request.NewPassword)
		if err != nil {
			return user.Core{}, errors.New("error while hashing new password")
		}

		log.Printf("berhasil hashing new password: %s", hash)
		request.Password = hash
	}

	if request.Name == "" && request.Email == "" && request.Password == "" {
		return user.Core{}, errors.New("failed to process the request due to empty input")
	}

	result, err := s.query.UpdateProfile(userId, request)
	if err != nil {
		if strings.Contains(err.Error(), "error while retrieving user profile") {
			return user.Core{}, errors.New("error while retrieving user profile")
		}
		if strings.Contains(err.Error(), "failed to update user, duplicate data entry") {
			return user.Core{}, errors.New("failed to update user, duplicate data entry")
		}
		return user.Core{}, errors.New("internal server error, failed to update account")
	}

	return result, nil
}

func (s *Service) Deactive(userId uint) (user.Core, error) {
	result, err := s.query.Deactive(userId)
	if err != nil {
		if strings.Contains("error while retrieving user profile", err.Error()) {
			return user.Core{}, errors.New("account not registered")
		}
		return user.Core{}, errors.New("internal server error, failed to delete account")
	}

	return result, nil
}

// GetAllUsers implements user.UserService
func (s *Service) GetAllUserHasBooks() ([]user.Core, error) {
	result, err := s.query.GetAllUserHasBooks()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return []user.Core{}, errors.New("not found, error while retrieving list users")
		} else {
			return []user.Core{}, errors.New("internal server error")
		}
	}
	return result, nil
}
