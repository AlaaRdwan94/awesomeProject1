package middlewares

import (
	"errors"
	"net/http"

	"github.com/AlaaRadwan94/awesomeProject1/api/responses"

	"github.com/AlaaRadwan94/awesomeProject1/api/auth"
)

//this function will format all responses to JSON
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//this function will check for the validity of the authentication token provided
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
