package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWToken(username string) (string, error) {
	secretKey := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
