package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/interface/api/dtos"
	"net/http"
)

type getAddressByUserId struct {
	addressRepository interfaces.GetAddressByUserId
}

func NewGetAddressByUserId(repo interfaces.GetAddressByUserId) *getAddressByUserId {
	return &getAddressByUserId{addressRepository: repo}
}
func (s *getAddressByUserId) GetByUserID(userId uint64) (addressesDtos []dtos.AddressDto, err *errors.Error) {

	addressesDtos, error := s.addressRepository.GetByUserID(userId)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	return
}
