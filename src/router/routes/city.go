package routes

import (
	"api/src/handlers"
	"net/http"
)

var cityRoutes = []Route{
	{
		Uri:     "/city",
		Method:  http.MethodGet,
		Handler: handlers.GetCities,
		ReqAuth: false,
	},
	{
		Uri:     "/city/{state_id}",
		Method:  http.MethodGet,
		Handler: handlers.GetCitiesByStateID,
		ReqAuth: false,
	},
}
