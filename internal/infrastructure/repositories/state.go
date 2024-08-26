package repositories

import (
	"database/sql"
	"taskmanager/internal/domain/entities"
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
func (s state) GetAll() ([]entities.State, error) {

	rows, err := s.db.Query("select id, name, uf from state")
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

// GetByID get a state by ID
func (s state) GetByID(id uint64) (entities.State, error) {

	rows, err := s.db.Query("select id, name, uf from state where id = $1", id)
	if err != nil {
		return entities.State{}, err
	}
	defer rows.Close()

	var state entities.State

	if rows.Next() {
		if err = rows.Scan(&state.ID, &state.Name, &state.Uf); err != nil {
			return entities.State{}, err
		}
	}

	return state, nil
}
