package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = Route{
	Uri:     "/auth",
	Method:  http.MethodPost,
	Handler: controllers.Auth,
	ReqAuth: false,
}
