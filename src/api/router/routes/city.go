package routes

import (
	"api/src/api/controllers"
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
		Uri:     "/city/{city_id}",
		Method:  http.MethodGet,
		Handler: controllers.GetCityByID,
		ReqAuth: false,
	},
	{
		Uri:     "/state/{state_id}/city",
		Method:  http.MethodGet,
		Handler: controllers.GetCitiesByStateID,
		ReqAuth: false,
	},
}
