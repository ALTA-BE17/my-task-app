package middlewares

import (
	"errors"
	"log"
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(dep dependency.Dependency) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWTSecret),
		SigningMethod: "HS256",
	})
}

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return "", nil
	}
	return result, err
}

func ExtractToken(e echo.Context) (uint, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := uint(claims["id"].(float64))
		return id, nil
	}
	return 0, errors.New("failed to extract jwt-token")
}
