package controllers

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/common/request"
	"api/src/application/common/responses"
	"api/src/application/common/security"
	"api/src/infrastructure/database"
	"api/src/infrastructure/repositories"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var auth dtos.Auth
	if err := request.ProcessBody(r, &auth); err != nil {
		responses.Error(w, err)
		return
	}

	db, err := database.NewDatabase()
	if err != nil {
		err := errors.NewError(err.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	user, err := userRepository.GetByEmail(auth.Email)
	if err != nil {
		responses.Error(w, errors.NewError("invalid credentials", http.StatusUnauthorized))
		return
	}

	if err = security.Compare(user.Password, auth.Password); err != nil {
		responses.Error(w, errors.NewError("invalid credentials", http.StatusUnauthorized))
		return
	}

	token, error := security.GenerateToken(user.ID)
	if error != nil {
		responses.Error(w, error)
		return
	}

	responses.Json(w, http.StatusOK, token)
}
