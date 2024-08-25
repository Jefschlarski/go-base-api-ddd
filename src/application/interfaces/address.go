package interfaces

import (
	"api/src/api/dtos"
	"api/src/domain/entities"
)

type CreateAddress interface {
	Create(dtos.CreateAddressDto) (uint64, error)
}

type DeleteAddress interface {
	Delete(id uint64) (int64, error)
}

type GetAddressByUserId interface {
	GetByUserID(uint64) ([]entities.Address, error)
}

type GetAddress interface {
	Get(uint64) (entities.Address, error)
}

type GetAllAddresses interface {
	GetAll() ([]entities.Address, error)
}

type UpdateAddress interface {
	Update(dtos.AddressDto) (int64, error)
}
