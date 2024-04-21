package handlers

import (
	"api/src/common/responses"
	"api/src/database"
	"api/src/repositories"
	"net/http"
)

// GetStates gets all states
func GetStates(w http.ResponseWriter, r *http.Request) {

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewStateRepository(db)

	states, err := repository.GetAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, states)
}
