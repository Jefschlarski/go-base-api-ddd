package routes

import (
	"api/src/controllers"
	"net/http"
)

var stateRoutes = Route{
	Uri:     "/state",
	Method:  http.MethodGet,
	Handler: controllers.GetStates,
	ReqAuth: false,
}
