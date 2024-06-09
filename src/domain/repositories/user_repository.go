package repositories

import (
	"api/src/interface/api/dtos"
)

type UserRepository interface {
	Create(dtos.CreateAddressDto) (uint64, error)
	GetAddressesByUserID(uint64) ([]dtos.AddressDto, error)
	UpdateAddressesByID(dtos.AddressDto) (int64, error)
	GetAddressById(uint64) (dtos.AddressDto, error)
	GetAddresses() ([]dtos.AddressDto, error)
	DeleteAddressByID(uint64) (int64, error)
}
