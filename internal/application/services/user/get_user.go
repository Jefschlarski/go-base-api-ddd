package services

import (
	"net/http"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

type getUser struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewGetUser(repo repositories.UserRepositoryInterface) interfaces.GetUser {
	return &getUser{UserRepository: repo}
}

func (s *getUser) Execute(id uint64) (entities.User, *errors.Error) {

	user, err := s.UserRepository.Get(id)
	if err != nil {
		return entities.User{}, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	return user, nil
}
