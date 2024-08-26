package repositories

import (
	"taskmanager/internal/api/dtos"
	"taskmanager/internal/domain/entities"
)

type AddressRepositoryInterface interface {
	Create(dtos.CreateAddressDto) (uint64, error)
	GetByUserID(uint64) ([]entities.Address, error)
	Update(dtos.AddressDto) (int64, error)
	Get(uint64) (entities.Address, error)
	GetAll() ([]entities.Address, error)
	Delete(id uint64) (int64, error)
}
