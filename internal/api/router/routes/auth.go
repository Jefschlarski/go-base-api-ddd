package routes

import (
	"net/http"
	"taskmanager/internal/api/controllers"
)

var authRoutes = Route{
	Uri:     "/auth",
	Method:  http.MethodPost,
	Handler: controllers.Auth,
	ReqAuth: false,
}
