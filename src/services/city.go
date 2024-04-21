package services

import (
	"api/src/common/errors"
	"api/src/database"
	"api/src/entities"
	"api/src/repositories"
	"net/http"
)

type cityService struct{}

func NewCityService() *cityService {
	return &cityService{}
}

func (s *cityService) GetAll() ([]entities.City, *errors.Error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository := repositories.NewCityRepository(db)

	cities, error := repository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return cities, nil
}

func (s *cityService) GetByStateID(id uint64) ([]entities.City, *errors.Error) {
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository := repositories.NewCityRepository(db)

	cities, error := repository.GetByStateID(id)
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return cities, nil
}
