package controllers

import (
	"net/http"
	cityServices "taskmanager/internal/application/services/city"
	"taskmanager/internal/common/request"
	"taskmanager/internal/common/responses"
	"taskmanager/internal/infrastructure/database"
	"taskmanager/internal/infrastructure/repositories"
)

// GetCities gets all states
func GetCities(w http.ResponseWriter, r *http.Request) {

	db := database.GetPostgresDB()

	getAllCities := cityServices.NewGetAllCities(repositories.NewCityRepository(db))

	cities, err := getAllCities.Execute()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}

func GetCityByID(w http.ResponseWriter, r *http.Request) {

	city_id, err := request.GetId(r, "city_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	db := database.GetPostgresDB()

	services := cityServices.NewGetCity(repositories.NewCityRepository(db))

	city, err := services.Get(city_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, city)
}

// GetCitiesByStateID gets all cities by state ID
func GetCitiesByStateID(w http.ResponseWriter, r *http.Request) {

	state_id, err := request.GetId(r, "state_id")
	print(state_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	db := database.GetPostgresDB()

	services := cityServices.NewGetCitiesByStateID(repositories.NewCityRepository(db))

	cities, err := services.GetByStateID(state_id)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}
