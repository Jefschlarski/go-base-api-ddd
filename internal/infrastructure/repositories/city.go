package repositories

import (
	"database/sql"
	"taskmanager/internal/domain/repositories"
	"taskmanager/internal/interface/dtos"
)

// city struct represents a city repository
type cityRepository struct {
	db *sql.DB
}

// NewCityRepository create a new city repository
func NewCityRepository(db *sql.DB) repositories.CityRepositoryInterface {
	return &cityRepository{db}
}

// GetAll get all cities in the database
func (c cityRepository) GetAll() (citiesDtoList []dtos.CityDto, err error) {

	rows, err := c.db.Query("select id, state_id, name from city")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var city dtos.CityDto
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return
		}
		citiesDtoList = append(citiesDtoList, city)
	}

	return
}

// GetByStateID get all cities by state ID
func (c cityRepository) GetByStateID(stateID uint64) (citiesDtoList []dtos.CityDto, err error) {

	rows, err := c.db.Query("select id, state_id, name from city where state_id = $1", stateID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var city dtos.CityDto
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return
		}
		citiesDtoList = append(citiesDtoList, city)
	}

	return
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
