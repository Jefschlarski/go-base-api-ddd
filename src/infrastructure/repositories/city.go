package repositories

import (
	repositoriesInterfaces "api/src/domain/repositories"
	"api/src/infrastructure/database"
	"api/src/interface/api/dtos"
)

// city struct represents a city repository
type cityRepository struct{}

// NewCityRepository create a new city repository
func NewCityRepository() repositoriesInterfaces.CityRepository {
	return &cityRepository{}
}

// GetAll get all cities in the database
func (u cityRepository) GetAll() ([]dtos.CityDto, error) {

	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, state_id, name from city")
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
func (u cityRepository) GetByStateID(stateID uint64) ([]dtos.CityDto, error) {

	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, state_id, name from city where state_id = $1", stateID)
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
func (u cityRepository) Get(id uint64) (city dtos.CityDto, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, state_id, name from city where id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	print(rows)

	if rows.Next() {
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return
		}
	}

	return
}
