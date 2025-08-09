package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey=[]byte("secret")

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"email": email,
	"exp":time.Now().Add(time.Hour*72).Unix(),
	})

	return token.SignedString(jwtKey)
}