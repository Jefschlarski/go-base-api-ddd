package controllers

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/common/request"
	"api/src/application/common/responses"
	"api/src/application/common/security"
	services "api/src/application/services/user"
	"api/src/domain/entities"
	"api/src/infrastructure/database"
	"api/src/infrastructure/repositories"
	"fmt"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user entities.User
	if err := request.ProcessBody(r, &user); err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	createUser := services.NewCreateUser(repositories.NewUserRepository(db))

	id, err := createUser.Execute(&user)
	if err != nil {
		responses.Error(w, err)
		return
	}

	user.ID = id

	responses.Json(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	getAllUsers := services.NewGetAllUsers(repositories.NewUserRepository(db))

	users, err := getAllUsers.Execute()
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	getUser := services.NewGetUser(repositories.NewUserRepository(db))

	user, err := getUser.Execute(userID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	if err := security.VerifyId(userID, r); err != nil {
		responses.Error(w, err)
		return
	}

	var user entities.User
	if err = request.ProcessBody(r, &user); err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	updateUser := services.NewUpdateUser(repositories.NewUserRepository(db))

	rowsAffected, err := updateUser.Execute(userID, user)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	if err = security.VerifyId(userID, r); err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	deleteUser := services.NewDeleteUser(repositories.NewUserRepository(db))

	rowsAffected, err := deleteUser.Execute(userID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Deleted %d rows", rowsAffected))
}

// UpdateUserPassword updates a user password
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	if err = security.VerifyId(userID, r); err != nil {
		responses.Error(w, err)
		return
	}

	var updatePassword dtos.UpdatePassword
	if err = request.ProcessBody(r, &updatePassword); err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	updateUserPassword := services.NewUpdateUserPassword(repositories.NewUserRepository(db))

	rowsAffected, err := updateUserPassword.Execute(userID, updatePassword)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}
