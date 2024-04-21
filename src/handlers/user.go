package handlers

import (
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/common/security"
	"api/src/database"
	"api/src/dtos"
	"api/src/entities"
	"api/src/repositories"
	"errors"
	"fmt"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user entities.User
	if err := request.ProcessBody(r, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(true); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.GetAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
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

	repository := repositories.NewUserRepository(db)

	user, err := repository.Get(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = security.VerifyId(userID, r); err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}

	var user entities.User
	if err = request.ProcessBody(r, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(false); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	rowsAffected, err := repository.Update(userID, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		error := fmt.Errorf("there is no user with id %d", userID)
		responses.Error(w, http.StatusNotFound, error)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = security.VerifyId(userID, r); err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	rowsAffected, err := repository.Delete(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		error := fmt.Errorf("there is no user with id %d", userID)
		responses.Error(w, http.StatusNotFound, error)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Deleted %d rows", rowsAffected))
}

// UpdateUserPassword updates a user password
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = security.VerifyId(userID, r); err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}

	var updatePassword dtos.UpdatePassword
	if err = request.ProcessBody(r, &updatePassword); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	currentPassword, err := repository.GetPassword(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Compare(currentPassword, updatePassword.OldPassword); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New("the current password provided doesn't match with the password saved in database"))
		return
	}

	if err = security.Compare(currentPassword, updatePassword.NewPassword); err == nil {
		responses.Error(w, http.StatusBadRequest, errors.New("the new password provided should be different from the current password"))
		return
	}

	hash, err := security.Hash(updatePassword.NewPassword)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	rowsAffected, err := repository.UpdatePassword(userID, string(hash))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}
