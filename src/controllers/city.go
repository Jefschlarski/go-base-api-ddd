package controllers

import (
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/services"
	"net/http"
)

// GetCities gets all states
func GetCities(w http.ResponseWriter, r *http.Request) {

	services := services.NewCityService()

	cities, err := services.GetAll()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}

// GetCitiesByStateID gets all cities by state ID
func GetCitiesByStateID(w http.ResponseWriter, r *http.Request) {
	state_id, err := request.GetId(r, "state_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	services := services.NewCityService()

	cities, err := services.GetByStateID(state_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}
