package routes

import "net/http"

// Routes is a struct to define routes
type Routes struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	ReqAuth bool
}
