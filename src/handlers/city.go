package handlers

import (
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/database"
	"api/src/repositories"
	"net/http"
)

// GetCities gets all states
func GetCities(w http.ResponseWriter, r *http.Request) {

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCityRepository(db)

	cities, err := repository.GetAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}

// GetCitiesByStateID gets all cities by state ID
func GetCitiesByStateID(w http.ResponseWriter, r *http.Request) {
	state_id, err := request.GetId(r, "state_id")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCityRepository(db)

	cities, err := repository.GetByStateID(state_id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, cities)
}
