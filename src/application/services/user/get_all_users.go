package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/domain/entities"
	"api/src/infrastructure/repositories"
	"net/http"
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
