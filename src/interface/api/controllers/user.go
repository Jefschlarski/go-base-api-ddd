package controllers

import (
	"api/src/application/common/request"
	"api/src/application/common/responses"
	"api/src/application/common/security"
	"api/src/application/services"
	"api/src/domain/entities"
	"api/src/interface/api/dtos"
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

	service := services.NewUserService()

	err := service.Create(&user)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	service := services.NewUserService()

	users, err := service.GetAll()
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

	service := services.NewUserService()

	user, err := service.Get(userID)
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

	service := services.NewUserService()

	rowsAffected, err := service.Update(userID, user)
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

	service := services.NewUserService()

	rowsAffected, err := service.Delete(userID)
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

	service := services.NewUserService()

	rowsAffected, err := service.UpdatePassword(userID, updatePassword)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}
