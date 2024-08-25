package services

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"net/http"
)

type updateAddress struct {
	addressRepository interfaces.UpdateAddress
}

func NewUpdateAddress(repo interfaces.UpdateAddress) *updateAddress {
	return &updateAddress{addressRepository: repo}
}

func (s *updateAddress) Update(addressDto *dtos.AddressDto) (rowsAffected int64, error *errors.Error) {

	rowsAffected, err := s.addressRepository.Update(*addressDto)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}
