package routes

import (
	"api/src/interface/api/controllers"
	"net/http"
)

var stateRoutes = Route{
	Uri:     "/state",
	Method:  http.MethodGet,
	Handler: controllers.GetStates,
	ReqAuth: false,
}
