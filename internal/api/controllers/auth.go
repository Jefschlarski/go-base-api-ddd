package controllers

import (
	"net/http"
	"taskmanager/internal/api/dtos"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/common/request"
	"taskmanager/internal/common/responses"
	"taskmanager/internal/common/security"
	"taskmanager/internal/infrastructure/database"
	"taskmanager/internal/infrastructure/repositories"
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
