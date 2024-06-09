package controllers

import (
	"api/src/application/common/request"
	"api/src/application/common/responses"
	cityServices "api/src/application/services/city"
	"api/src/infrastructure/repositories"
	"fmt"
	"net/http"
)

// GetCities gets all states
func GetCities(w http.ResponseWriter, r *http.Request) {

	getAllCities := cityServices.NewGetAllCities(repositories.NewCityRepository())

	cities, err := getAllCities.Execute()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}

func GetCityByID(w http.ResponseWriter, r *http.Request) {

	city_id, err := request.GetId(r, "city_id")
	fmt.Println(city_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	services := cityServices.NewGetCity(repositories.NewCityRepository())

	city, err := services.Get(city_id)

	fmt.Println(city.ID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, city)
}

// GetCitiesByStateID gets all cities by state ID
func GetCitiesByStateID(w http.ResponseWriter, r *http.Request) {
	state_id, err := request.GetId(r, "state_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	services := cityServices.NewGetCitiesByStateID(repositories.NewCityRepository())

	cities, err := services.GetByStateID(state_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}
