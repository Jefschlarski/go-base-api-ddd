package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
)

type createAddress struct {
	addressRepository interfaces.CreateAddress
	cityRepository    interfaces.GetCity
}

func NewCreateAddress(addressRepo interfaces.CreateAddress, cityRepo interfaces.GetCity) *createAddress {
	return &createAddress{addressRepository: addressRepo, cityRepository: cityRepo}
}

func (s *createAddress) Create(createAddressDto *dtos.CreateAddressDto) *errors.Error {

	_, err := s.cityRepository.Get(createAddressDto.CityID)
	if err != nil {
		return errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	_, error := s.addressRepository.Create(*createAddressDto)
	if error != nil {
		return errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return nil
}
