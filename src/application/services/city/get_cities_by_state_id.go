package services

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"net/http"
)

type getCitiesByStateID struct {
	cityRepository interfaces.GetCitiesByStateID
}

func NewGetCitiesByStateID(repo interfaces.GetCitiesByStateID) *getCitiesByStateID {
	return &getCitiesByStateID{cityRepository: repo}
}

func (s *getCitiesByStateID) GetByStateID(id uint64) ([]dtos.CityDto, *errors.Error) {
	cities, error := s.cityRepository.GetByStateID(id)
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return cities, nil
}
