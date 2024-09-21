package controllers

import (
	"fmt"
	"net/http"
	userServices "taskmanager/internal/application/services/user"
	"taskmanager/internal/common/request"
	"taskmanager/internal/common/responses"
	"taskmanager/internal/common/security"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/infrastructure/pg"
	"taskmanager/internal/infrastructure/repositories"
	"taskmanager/internal/interface/dtos"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user entities.User
	if err := request.ProcessBody(r, &user); err != nil {
		responses.Error(w, err)
		return
	}

	db := pg.GetDB()
	createUser := userServices.NewCreateUser(repositories.NewUserRepository(db))

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

	db := pg.GetDB()

	getAllUsers := userServices.NewGetAllUsers(repositories.NewUserRepository(db))

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

	db := pg.GetDB()

	getUser := userServices.NewGetUser(repositories.NewUserRepository(db))

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

	db := pg.GetDB()

	updateUser := userServices.NewUpdateUser(repositories.NewUserRepository(db))

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

	db := pg.GetDB()

	deleteUser := userServices.NewDeleteUser(repositories.NewUserRepository(db))

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

	db := pg.GetDB()

	updateUserPassword := userServices.NewUpdateUserPassword(repositories.NewUserRepository(db))

	rowsAffected, err := updateUserPassword.Execute(userID, updatePassword)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}
