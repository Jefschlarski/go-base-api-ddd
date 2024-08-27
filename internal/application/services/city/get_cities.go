package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/interface/dtos"
)

type getAllCities struct {
	cityRepository interfaces.GetAllCities
}

func NewGetAllCities(repo interfaces.GetAllCities) *getAllCities {
	return &getAllCities{cityRepository: repo}
}

func (s *getAllCities) Execute() ([]dtos.CityDto, *errors.Error) {
	cities, error := s.cityRepository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return cities, nil
}
