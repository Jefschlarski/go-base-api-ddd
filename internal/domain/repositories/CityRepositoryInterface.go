package repositories

import "taskmanager/internal/api/dtos"

type CityRepositoryInterface interface {
	Get(uint64) (dtos.CityDto, error)
	GetAll() ([]dtos.CityDto, error)
	GetByStateID(uint64) ([]dtos.CityDto, error)
}
