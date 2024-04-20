package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes is a struct to define routes
type Routes struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	ReqAuth bool
}

// ConfigRoutes set all routes in the router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
