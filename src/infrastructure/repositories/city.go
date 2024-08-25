package repositories

import (
	"api/src/api/dtos"
	"api/src/infrastructure/database"
)

type CityRepositoryInterface interface {
	Get(uint64) (dtos.CityDto, error)
	GetAll() ([]dtos.CityDto, error)
	GetByStateID(uint64) ([]dtos.CityDto, error)
}

// city struct represents a city repository
type cityRepository struct {
	db database.DatabaseInterface
}

// NewCityRepository create a new city repository
func NewCityRepository(db database.DatabaseInterface) CityRepositoryInterface {
	return &cityRepository{db}
}

// GetAll get all cities in the database
func (c cityRepository) GetAll() ([]dtos.CityDto, error) {

	rows, err := c.db.Query("select id, state_id, name from city")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []dtos.CityDto

	for rows.Next() {
		var city dtos.CityDto
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

// GetByStateID get all cities by state ID
func (c cityRepository) GetByStateID(stateID uint64) ([]dtos.CityDto, error) {

	rows, err := c.db.Query("select id, state_id, name from city where state_id = $1", stateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []dtos.CityDto

	for rows.Next() {
		var city dtos.CityDto
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

// Get get a city by ID and relation
func (c cityRepository) Get(id uint64) (city dtos.CityDto, err error) {

	rows, err := c.db.Query("select id, state_id, name from city where id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return
		}
	}

	return
}
