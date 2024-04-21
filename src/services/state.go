package services

import (
	"api/src/common/errors"
	"api/src/database"
	"api/src/entities"
	"api/src/repositories"
	"net/http"
)

type stateService struct{}

func NewStateService() *stateService {
	return &stateService{}
}

func (s *stateService) GetAll() ([]entities.State, *errors.Error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository := repositories.NewStateRepository(db)

	states, error := repository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return states, nil
}
