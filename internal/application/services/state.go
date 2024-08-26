package services

import (
	"net/http"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/infrastructure/database"
	"taskmanager/internal/infrastructure/repositories"
)

type stateService struct{}

func NewStateService() *stateService {
	return &stateService{}
}

func (s *stateService) GetAll() ([]entities.State, *errors.Error) {
	db, error := database.OpenConnection()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	repository := repositories.NewStateRepository(db)

	states, error := repository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return states, nil
}
