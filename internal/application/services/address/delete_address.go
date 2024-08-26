package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
)

type deleteAddress struct {
	addressRepository interfaces.DeleteAddress
}

func NewDeleteAddress(repo interfaces.DeleteAddress) *deleteAddress {
	return &deleteAddress{addressRepository: repo}
}

func (s *deleteAddress) Delete(id uint64) (rowsAffected int64, error *errors.Error) {

	rowsAffected, err := s.addressRepository.Delete(id)
	if err != nil {
		error = errors.NewError(err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		error = errors.NewError("no rows affected", http.StatusBadRequest)
		return
	}

	return
}
