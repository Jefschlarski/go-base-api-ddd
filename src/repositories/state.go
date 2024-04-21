package repositories

import (
	"api/src/entities"
	"database/sql"
)

// state struct represents a state repository
type state struct {
	db *sql.DB
}

// NewStateRepository create a new state repository
func NewStateRepository(db *sql.DB) *state {
	return &state{db}
}

// GetAll get all states
func (u state) GetAll() ([]entities.State, error) {

	rows, err := u.db.Query("select id, name, uf from state")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var states []entities.State

	for rows.Next() {
		var state entities.State
		if err = rows.Scan(&state.ID, &state.Name, &state.Uf); err != nil {
			return nil, err
		}
		states = append(states, state)
	}

	return states, nil
}
