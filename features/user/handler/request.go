package handler

import "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type UpdateProfileRequest struct {
	Name              *string `json:"name" form:"name"`
	Email             *string `json:"email" form:"email"`
	Password          *string `json:"password" form:"password"`
	NewPassword       *string `json:"new_password" form:"new_password"`
	ConfirmedPassword *string `json:"confirmed_password" form:"confirmed_password"`
}

func RequestToCore(data interface{}) user.Core {
	res := user.Core{}
	switch v := data.(type) {
	case RegisterRequest:
		res.Name = v.Name
		res.Email = v.Email
		res.Password = v.Password
	case LoginRequest:
		res.Name = v.Name
		res.Password = v.Password
	case *UpdateProfileRequest:
		if v.Name != nil {
			res.Name = *v.Name
		}
		if v.Email != nil {
			res.Email = *v.Email
		}
		if v.Password != nil {
			res.Password = *v.Password
		}
		if v.NewPassword != nil {
			res.NewPassword = *v.NewPassword
		}
		if v.ConfirmedPassword != nil {
			res.ConfirmedPassword = *v.ConfirmedPassword
		}
	default:
		return user.Core{}
	}
	return res
}
