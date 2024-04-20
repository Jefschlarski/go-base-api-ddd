package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenRouter returns a new router
func GenRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
