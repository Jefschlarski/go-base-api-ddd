package router

import "github.com/gorilla/mux"

// GenRouter returns a new router
func GenRouter() *mux.Router {
	return mux.NewRouter()
}
