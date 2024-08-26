package routes

import (
	"net/http"
	"taskmanager/internal/api/controllers"
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
