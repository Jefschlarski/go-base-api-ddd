package repositories

import (
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/interface/dtos"
)

type AddressRepositoryInterface interface {
	Create(addressDto dtos.CreateAddressDto) (id uint64, err error)
	GetByUserID(id uint64) (addressesList []entities.Address, err error)
	Update(addressDto dtos.AddressDto) (rowsAffected int64, err error)
	Get(id uint64) (address entities.Address, err error)
	GetAll() (addressesList []entities.Address, err error)
	Delete(id uint64) (rowsAffected int64, err error)
}
