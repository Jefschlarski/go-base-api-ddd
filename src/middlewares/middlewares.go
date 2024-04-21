package middlewares

import (
	"api/src/common/responses"
	"api/src/common/security"
	"fmt"
	"net/http"
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
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
