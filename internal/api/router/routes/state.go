package routes

import (
	"net/http"
	"taskmanager/internal/api/controllers"
)

var stateRoutes = Route{
	Uri:     "/state",
	Method:  http.MethodGet,
	Handler: controllers.GetStates,
	ReqAuth: false,
}
