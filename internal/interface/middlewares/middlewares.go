package middlewares

import (
	"fmt"
	"net/http"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/common/responses"
	"taskmanager/internal/common/security"
)

// Logger function is a middleware function that intercepts HTTP requests and logs the request details.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// authenticate function is a middleware function that intercepts HTTP requests and performs authentication logic before passing the request to the next handler.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			responses.Error(w, errors.NewError(err.Error(), http.StatusUnauthorized))
			return
		}
		next(w, r)
	}
}
