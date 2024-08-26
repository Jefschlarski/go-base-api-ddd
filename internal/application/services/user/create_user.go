package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

type createUser struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewCreateUser(repo repositories.UserRepositoryInterface) interfaces.CreateUser {
	return &createUser{UserRepository: repo}
}

func (s *createUser) Execute(user *entities.User) (uint64, *errors.Error) {

	if error := user.Prepare(true); error != nil {
		return 0, error
	}

	id, err := s.UserRepository.Create(user)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusBadRequest)
	}

	return id, nil
}
