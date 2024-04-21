package routes

import (
	"api/src/controllers"
	"net/http"
)

var cityRoutes = []Route{
	{
		Uri:     "/city",
		Method:  http.MethodGet,
		Handler: controllers.GetCities,
		ReqAuth: false,
	},
	{
		Uri:     "/city/{state_id}",
		Method:  http.MethodGet,
		Handler: controllers.GetCitiesByStateID,
		ReqAuth: false,
	},
}
