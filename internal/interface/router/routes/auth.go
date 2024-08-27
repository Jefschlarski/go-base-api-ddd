package routes

import (
	"net/http"
	"taskmanager/internal/interface/controllers"
)

var authRoutes = Route{
	Uri:     "/auth",
	Method:  http.MethodPost,
	Handler: controllers.Auth,
	ReqAuth: false,
}
