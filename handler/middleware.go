package handler

import (
	"net/http"
	"imageService/jwtValidator"
	"errors"
)

var (
	validationErr = errors.New("invalid token. Please register here http://localhost:8081/registration to get new token")
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			h.ServeHTTP(writer, request)
		} else {
			token := request.Header.Get("token")
			if valid := jwtValidator.Validate(token); valid {
				h.ServeHTTP(writer, request)
			} else {
				http.Error(writer, validationErr.Error(), 401)
				return
			}
		}
	})
}
