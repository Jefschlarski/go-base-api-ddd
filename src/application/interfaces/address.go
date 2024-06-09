package interfaces

import (
	"api/src/interface/api/dtos"
)

type CreateAddress interface {
	Create(dtos.CreateAddressDto) (uint64, error)
}

type DeleteAddress interface {
	Delete(id uint64) (int64, error)
}

type GetAddressByUserId interface {
	GetByUserID(uint64) ([]dtos.AddressDto, error)
}

type GetAddress interface {
	Get(uint64) (dtos.AddressDto, error)
}

type GetAllAddresses interface {
	GetAll() ([]dtos.AddressDto, error)
}

type UpdateAddress interface {
	Update(dtos.AddressDto) (int64, error)
}
