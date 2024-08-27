package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

type getAllStates struct {
	stateRepository repositories.StateRepositoryInterface
}

func NewGetAllStates(repo repositories.StateRepositoryInterface) interfaces.GetAllStates {
	return &getAllStates{stateRepository: repo}
}

func (g *getAllStates) Execute() ([]entities.State, *errors.Error) {
	states, error := g.stateRepository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return states, nil
}
