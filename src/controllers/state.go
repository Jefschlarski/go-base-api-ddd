package controllers

import (
	"api/src/common/responses"
	"api/src/services"
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
