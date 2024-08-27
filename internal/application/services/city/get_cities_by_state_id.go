package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/interface/dtos"
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
