package router

import (
	"taskmanager/internal/api/router/routes"

	"github.com/gorilla/mux"
)

// GenRouter returns a new router
func GenRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
