package controllers

import (
	"api/src/application/common/responses"
	"api/src/application/services"
	"net/http"
)

// GetStates gets all states
func GetStates(w http.ResponseWriter, r *http.Request) {

	services := services.NewStateService()

	states, err := services.GetAll()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, states)
}
