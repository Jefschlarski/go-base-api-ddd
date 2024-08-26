package controllers

import (
	"net/http"
	"taskmanager/internal/application/services"
	"taskmanager/internal/common/responses"
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
