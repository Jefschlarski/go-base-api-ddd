package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
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
