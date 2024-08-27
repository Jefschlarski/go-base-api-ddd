package repositories

import (
	"taskmanager/internal/domain/entities"
)

type StateRepositoryInterface interface {
	GetAll() (statesList []entities.State, err error)
}
