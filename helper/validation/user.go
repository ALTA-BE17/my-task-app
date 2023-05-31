package validation

import (
	"errors"
	"fmt"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/go-playground/validator/v10"
	pv "github.com/wagslane/go-password-validator"
)

type Register struct {
	Name     string `validate:"required"`
	Password string `validate:"required,passwordValidator"`
}

type Login struct {
	Name     string
	Password string
}

func UserValidate(option string, data interface{}) interface{} {
	switch option {
	case "register":
		res := Register{}
		if v, ok := data.(user.Core); ok {
			res.Name = v.Name
			res.Password = v.Password
		}
		return res
	case "login":
		res := Login{}
		if v, ok := data.(Login); ok {
			res.Name = v.Name
			res.Password = v.Password
		}
		return res
	default:
		return nil
	}
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	entropy := pv.GetEntropy(password)
	fmt.Printf("Entropy: %.2f bits\n", entropy)

	const minEntropyBits = 60
	err := pv.Validate(password, minEntropyBits)

	return err == nil
}

func Authenticate(data interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("passwordValidator", PasswordValidator)
	if err != nil {
		return err
	}

	err1 := validate.Struct(data)
	if err1 != nil {
		return err1
	}
	return nil
}

func UpdatePasswordValidator(password string) error {
	entropy := pv.GetEntropy(password)
	fmt.Printf("Entropy: %.2f bits\n", entropy)

	const minEntropyBits = 60
	err := pv.Validate(password, minEntropyBits)

	if err != nil {
		return errors.New("password strength is low")
	}

	return nil
}
