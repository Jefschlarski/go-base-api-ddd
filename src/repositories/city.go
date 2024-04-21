package repositories

import (
	"api/src/entities"
	"database/sql"
)

// city struct represents a city repository
type city struct {
	db *sql.DB
}

// NewCityRepository create a new city repository
func NewCityRepository(db *sql.DB) *city {
	return &city{db}
}

// GetAll get all cities in the database
func (u city) GetAll() ([]entities.City, error) {

	rows, err := u.db.Query("select id, state_id, name from city")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []entities.City

	for rows.Next() {
		var city entities.City
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

// GetByStateID get all cities by state ID
func (u city) GetByStateID(stateID uint64) ([]entities.City, error) {

	rows, err := u.db.Query("select id, state_id, name from city where state_id = $1", stateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []entities.City

	for rows.Next() {
		var city entities.City
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

// GetByID get a city by ID
func (u city) GetByID(id uint64, relation bool) (entities.City, error) {

	rows, err := u.db.Query("select id, state_id, name from city where id = $1", id)
	if err != nil {
		return entities.City{}, err
	}
	defer rows.Close()

	var city entities.City

	if rows.Next() {
		if err = rows.Scan(&city.ID, &city.StateID, &city.Name); err != nil {
			return entities.City{}, err
		}
	}

	if relation {
		stateRepository := NewStateRepository(u.db)
		state, err := stateRepository.GetByID(city.StateID)
		if err != nil {
			return entities.City{}, err
		}
		city.State = state
	}

	return city, nil
}
