package handler

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
)

type SearchUsersResponse struct {
	Username  string           `json:"username"`
	Phone     string           `json:"phone"`
	Email     string           `json:"email"`
	CreatedAt helper.LocalTime `json:"created_at"`
	UpdatedAt helper.LocalTime `json:"updated_at"`
}

func SearchUsers(u user.Core) SearchUsersResponse {
	return SearchUsersResponse{
		Username:  u.Username,
		Phone:     u.Phone,
		Email:     u.Email,
		CreatedAt: helper.LocalTime(u.CreatedAt),
		UpdatedAt: helper.LocalTime(u.UpdatedAt),
	}
}

type ProfileResponse struct {
	Username  string           `json:"username"`
	Phone     string           `json:"phone"`
	Email     string           `json:"email"`
	CreatedAt helper.LocalTime `json:"created_at"`
	UpdatedAt helper.LocalTime `json:"updated_at"`
}

func Profile(u user.Core) ProfileResponse {
	return ProfileResponse{
		Username:  u.Username,
		Phone:     u.Phone,
		Email:     u.Email,
		CreatedAt: helper.LocalTime(u.CreatedAt),
		UpdatedAt: helper.LocalTime(u.UpdatedAt),
	}
}
