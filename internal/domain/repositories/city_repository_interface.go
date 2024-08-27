package repositories

import "taskmanager/internal/interface/dtos"

type CityRepositoryInterface interface {
	Get(id uint64) (cityDto dtos.CityDto, err error)
	GetAll() (citiesDtoList []dtos.CityDto, err error)
	GetByStateID(id uint64) (citiesDtoList []dtos.CityDto, err error)
}
