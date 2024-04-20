package routes

import (
	"api/src/handlers"
	"net/http"
)

var authRoutes = Route{
	Uri:     "/auth",
	Method:  http.MethodPost,
	Handler: handlers.Auth,
	ReqAuth: false,
}
