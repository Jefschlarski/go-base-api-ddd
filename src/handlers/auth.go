package handlers

import (
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/common/security"
	"api/src/database"
	"api/src/entities"
	"api/src/repositories"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var auth entities.Auth
	if err := request.ProcessBody(r, &auth); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)

	user, err := userRepository.GetByEmail(auth.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Compare(user.Password, auth.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	responses.Json(w, http.StatusOK, user)
}
