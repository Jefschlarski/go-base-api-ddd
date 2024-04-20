package handlers

import (
	"api/src/common/responses"
	"api/src/database"
	"api/src/entities"
	"api/src/repositories"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user entities.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
