package repositories

import (
	"api/src/application/interfaces"
)

type CityRepository interface {
	interfaces.GetCity
	interfaces.GetAllCities
	interfaces.GetCitiesByStateID
}
