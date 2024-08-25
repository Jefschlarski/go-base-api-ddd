package routes

import (
	"api/src/api/controllers"
	"net/http"
)

var authRoutes = Route{
	Uri:     "/auth",
	Method:  http.MethodPost,
	Handler: controllers.Auth,
	ReqAuth: false,
}
