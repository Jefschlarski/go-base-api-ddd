package routes

import (
	"api/src/api/controllers"
	"net/http"
)

var stateRoutes = Route{
	Uri:     "/state",
	Method:  http.MethodGet,
	Handler: controllers.GetStates,
	ReqAuth: false,
}
