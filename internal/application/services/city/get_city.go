package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/interface/dtos"
)

type getCity struct {
	cityRepository interfaces.GetCity
}

func NewGetCity(repo interfaces.GetCity) *getCity {
	return &getCity{cityRepository: repo}
}
func (s *getCity) Get(id uint64) (cityDto dtos.CityDto, err *errors.Error) {

	cityDto, error := s.cityRepository.Get(id)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	if (cityDto == dtos.CityDto{}) {
		err = errors.NewError("city not found", http.StatusNotFound)
		return
	}

	return
}
