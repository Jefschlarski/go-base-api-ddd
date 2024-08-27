package controllers

import (
	"net/http"
	stateServices "taskmanager/internal/application/services/state"
	"taskmanager/internal/common/responses"
	"taskmanager/internal/infrastructure/database"
	"taskmanager/internal/infrastructure/repositories"
)

// GetStates gets all states
func GetStates(w http.ResponseWriter, r *http.Request) {
	db := database.GetPostgresDB()

	getAllStates := stateServices.NewGetAllStates(repositories.NewStateRepository(db))
	states, err := getAllStates.Execute()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, states)
}
