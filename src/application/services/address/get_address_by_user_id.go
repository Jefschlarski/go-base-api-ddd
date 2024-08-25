package services

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"net/http"
)

type getAddressByUserId struct {
	addressRepository interfaces.GetAddressByUserId
	cityRepository    interfaces.GetCity
}

func NewGetAddressByUserId(repo interfaces.GetAddressByUserId, cityRepo interfaces.GetCity) *getAddressByUserId {
	return &getAddressByUserId{addressRepository: repo, cityRepository: cityRepo}
}
func (s *getAddressByUserId) GetByUserID(userId uint64) (addressesDtos []dtos.ReturnAddressDto, err *errors.Error) {

	addresses, error := s.addressRepository.GetByUserID(userId)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		return
	}

	for _, address := range addresses {
		city, error := s.cityRepository.Get(address.CityID)
		if error != nil {
			err = errors.NewError(error.Error(), http.StatusInternalServerError)
			return
		}

		addressesDtos = append(addressesDtos, dtos.ReturnAddressDto{
			ID:         address.ID,
			UserID:     address.UserID,
			Complement: address.Complement,
			Number:     address.Number,
			Cep:        address.Cep,
			City:       city,
			CreatedAt:  address.CreatedAt,
		})
	}

	return
}
