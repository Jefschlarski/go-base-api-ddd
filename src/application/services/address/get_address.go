package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
)

type getAddress struct {
	addressRepository interfaces.GetAddress
}

func NewGetAddress(repo interfaces.GetAddress) *getAddress {
	return &getAddress{addressRepository: repo}
}

func (s *getAddress) Get(id uint64) (addressDto dtos.AddressDto, err *errors.Error) {

	addressDto, error := s.addressRepository.Get(id)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	if addressDto == (dtos.AddressDto{}) {
		err = errors.NewError("address not found", http.StatusNotFound)
	}

	return
}
