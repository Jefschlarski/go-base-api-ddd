package routes

import (
	"api/src/handlers"
	"net/http"
)

var stateRoutes = Route{
	Uri:     "/state",
	Method:  http.MethodGet,
	Handler: handlers.GetStates,
	ReqAuth: false,
}
