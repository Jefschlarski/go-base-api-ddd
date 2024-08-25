package services

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"net/http"
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
