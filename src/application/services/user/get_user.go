package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/domain/entities"
	"api/src/infrastructure/repositories"
	"net/http"
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
