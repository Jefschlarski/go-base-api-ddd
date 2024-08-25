package interfaces

import (
	"api/src/api/dtos"
)

type GetCity interface {
	Get(uint64) (dtos.CityDto, error)
}

type GetAllCities interface {
	GetAll() ([]dtos.CityDto, error)
}

type GetCitiesByStateID interface {
	GetByStateID(uint64) ([]dtos.CityDto, error)
}
