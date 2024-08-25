package services

import (
	"api/src/application/common/errors"
	"api/src/application/interfaces"
	"api/src/infrastructure/repositories"
	"net/http"
)

type deleteUser struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewDeleteUser(repo repositories.UserRepositoryInterface) interfaces.DeleteUser {
	return &deleteUser{UserRepository: repo}
}

func (s *deleteUser) Execute(id uint64) (rowsAffected int64, error *errors.Error) {

	rowsAffected, err := s.UserRepository.Delete(id)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}
