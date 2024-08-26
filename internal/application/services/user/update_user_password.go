package services

import (
	"net/http"
	"taskmanager/internal/api/dtos"
	"taskmanager/internal/application/interfaces"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/common/security"
	"taskmanager/internal/domain/repositories"
)

type updateUserPassword struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUpdateUserPassword(repo repositories.UserRepositoryInterface) interfaces.UpdateUserPassword {
	return &updateUserPassword{UserRepository: repo}
}

func (s *updateUserPassword) Execute(id uint64, updatePasswordDTO dtos.UpdatePassword) (rowsAffected int64, error *errors.Error) {

	currentPassword, err := s.UserRepository.GetPassword(id)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if err = security.Compare(currentPassword, updatePasswordDTO.OldPassword); err != nil {
		return 0, errors.NewError("the current password provided doesn't match with the password saved in database", http.StatusBadRequest)
	}

	if err = security.Compare(currentPassword, updatePasswordDTO.NewPassword); err == nil {
		return 0, errors.NewError("the new password provided should be different from the current password", http.StatusBadRequest)
	}

	hash, err := security.Hash(updatePasswordDTO.NewPassword)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, err = s.UserRepository.UpdatePassword(id, string(hash))
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}
