package jwtValidator

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

var key = []byte("DataArt")

func Validate(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return false
	}
	_, ok := token.Claims.(jwt.MapClaims)
	return ok && token.Valid
}
