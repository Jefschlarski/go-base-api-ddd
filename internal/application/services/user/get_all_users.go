package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

type getAllUsers struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewGetAllUsers(repo repositories.UserRepositoryInterface) interfaces.GetAllUsers {
	return &getAllUsers{UserRepository: repo}
}

func (s *getAllUsers) Execute() ([]entities.User, *errors.Error) {

	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	return users, nil
}
