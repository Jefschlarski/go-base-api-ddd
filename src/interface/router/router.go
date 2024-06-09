package router

import (
	"api/src/interface/router/routes"

	"github.com/gorilla/mux"
)

// GenRouter returns a new router
func GenRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
