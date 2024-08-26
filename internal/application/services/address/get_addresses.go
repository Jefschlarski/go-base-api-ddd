package services

import (
	"net/http"
	"taskmanager/internal/api/dtos"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
)

type getAllAddresses struct {
	addressRepository interfaces.GetAllAddresses
}

func NewGetAllAddresses(repo interfaces.GetAllAddresses) *getAllAddresses {
	return &getAllAddresses{addressRepository: repo}
}

func (s *getAllAddresses) GetAll() (addressesDto []dtos.AddressDto, err *errors.Error) {

	addresses, error := s.addressRepository.GetAll()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	for _, address := range addresses {
		addressesDto = append(addressesDto, dtos.AddressDto{
			ID:         address.ID,
			UserID:     address.UserID,
			Complement: address.Complement,
			Number:     address.Number,
			Cep:        address.Cep,
			CityID:     address.CityID,
		})
	}

	return
}
