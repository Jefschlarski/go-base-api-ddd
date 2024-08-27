package services

import (
	"fmt"
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/interface/dtos"
)

type createAddress struct {
	addressRepository interfaces.CreateAddress
	cityRepository    interfaces.GetCity
}

func NewCreateAddress(addressRepo interfaces.CreateAddress, cityRepo interfaces.GetCity) *createAddress {
	return &createAddress{addressRepository: addressRepo, cityRepository: cityRepo}
}

func (s *createAddress) Create(createAddressDto *dtos.CreateAddressDto) *errors.Error {

	city, err := s.cityRepository.Get(createAddressDto.CityID)
	if err != nil {
		return errors.NewError(err.Error(), http.StatusInternalServerError)
	}
	if city == (dtos.CityDto{}) {
		return errors.NewError(fmt.Sprintf("city with id %d not found", createAddressDto.CityID), http.StatusNotFound)
	}

	_, error := s.addressRepository.Create(*createAddressDto)
	if error != nil {
		return errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return nil
}
