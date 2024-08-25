package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/domain/entities"
	"api/src/infrastructure/repositories"
	"net/http"
)

type updateUser struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUpdateUser(repo repositories.UserRepositoryInterface) interfaces.UpdateUser {
	return &updateUser{UserRepository: repo}
}

func (s *updateUser) Execute(id uint64, user entities.User) (rowsAffected int64, error *errors.Error) {
	if error = user.Prepare(false); error != nil {
		return
	}

	rowsAffected, err := s.UserRepository.Update(id, user)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}
