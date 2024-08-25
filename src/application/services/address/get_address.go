package services

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"net/http"
)

type getAddress struct {
	addressRepository interfaces.GetAddress
}

func NewGetAddress(repo interfaces.GetAddress) *getAddress {
	return &getAddress{addressRepository: repo}
}

func (s *getAddress) Get(id uint64) (addressDto dtos.AddressDto, err *errors.Error) {

	address, error := s.addressRepository.Get(id)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	addressDto = dtos.AddressDto{
		ID:         address.ID,
		UserID:     address.UserID,
		Complement: address.Complement,
		Number:     address.Number,
		Cep:        address.Cep,
		CityID:     address.CityID,
	}

	if addressDto == (dtos.AddressDto{}) {
		err = errors.NewError("address not found", http.StatusNotFound)
	}

	return
}
