package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct to define routes
type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	ReqAuth bool
}

// ConfigRoutes set all routes in the router
func ConfigRoutes(r *mux.Router) *mux.Router {

	routes := userRoutes

	routes = append(routes, authRoutes)

	routes = append(routes, stateRoutes)

	routes = append(routes, cityRoutes...)

	for _, route := range routes {

		if route.ReqAuth {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Handler)).Methods(route.Method)
		}

	}
	return r
}
