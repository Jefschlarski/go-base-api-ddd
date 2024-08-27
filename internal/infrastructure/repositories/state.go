package repositories

import (
	"database/sql"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

// state struct represents a state repository
type stateRepository struct {
	db *sql.DB
}

// NewStateRepository create a new state repository
func NewStateRepository(db *sql.DB) repositories.StateRepositoryInterface {
	return &stateRepository{db}
}

// GetAll get all states
func (s stateRepository) GetAll() (statesList []entities.State, err error) {

	rows, err := s.db.Query("select id, name, uf from state")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var state entities.State
		if err = rows.Scan(&state.ID, &state.Name, &state.Uf); err != nil {
			return
		}
		statesList = append(statesList, state)
	}

	return
}

// GetByID get a state by ID
func (s stateRepository) GetByID(id uint64) (state entities.State, err error) {

	rows, err := s.db.Query("select id, name, uf from state where id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&state.ID, &state.Name, &state.Uf); err != nil {
			return
		}
	}

	return
}
