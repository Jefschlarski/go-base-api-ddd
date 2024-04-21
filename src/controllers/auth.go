package controllers

import (
	"api/src/common/errors"
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/common/security"
	"api/src/database"
	"api/src/dtos"
	"api/src/repositories"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var auth dtos.Auth
	if err := request.ProcessBody(r, &auth); err != nil {
		responses.Error(w, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	user, error := userRepository.GetByEmail(auth.Email)
	if error != nil {
		responses.Error(w, errors.NewError("invalid credentials", http.StatusUnauthorized))
		return
	}

	if error = security.Compare(user.Password, auth.Password); error != nil {
		responses.Error(w, errors.NewError("invalid credentials", http.StatusUnauthorized))
		return
	}

	token, err := security.GenerateToken(user.ID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, token)
}
