package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
)

type getCitiesByStateID struct {
	cityRepository interfaces.GetAllCities
}

func NewGetCitiesByStateID(repo interfaces.GetAllCities) *getCitiesByStateID {
	return &getCitiesByStateID{cityRepository: repo}
}

func (s *getCitiesByStateID) GetByStateID(id uint64) ([]dtos.CityDto, *errors.Error) {
	cities, error := s.cityRepository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return cities, nil
}
