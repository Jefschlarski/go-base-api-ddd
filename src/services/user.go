package services

import (
	"api/src/common/errors"
	"api/src/common/security"
	"api/src/database"
	"api/src/dtos"
	"api/src/entities"
	"api/src/repositories"
	"net/http"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) Create(user *entities.User) *errors.Error {
	if err := user.Prepare(true); err != nil {
		return err
	}

	db, err := database.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	id, error := repository.Create(*user)
	if error != nil {
		return errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	user.ID = id

	return nil
}

// GetAll gets all users
func (s *userService) GetAll() ([]entities.User, *errors.Error) {

	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, error := repository.GetAll()
	if error != nil {
		return nil, errors.NewError(error.Error(), http.StatusInternalServerError)
	}

	return users, nil
}

// Get gets a user
func (s *userService) Get(id uint64) (entities.User, *errors.Error) {
	db, err := database.OpenConnection()
	if err != nil {
		return entities.User{}, err
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, error := repository.Get(id)
	if error != nil {
		return entities.User{}, errors.NewError(error.Error(), http.StatusBadRequest)
	}

	return user, nil
}

// Update updates a user
func (s *userService) Update(id uint64, user entities.User) (rowsAffected int64, error *errors.Error) {
	if error = user.Prepare(false); error != nil {
		return
	}

	db, error := database.OpenConnection()
	if error != nil {
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	rowsAffected, err := repository.Update(id, user)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}

// Delete deletes a user
func (s *userService) Delete(id uint64) (rowsAffected int64, error *errors.Error) {
	db, error := database.OpenConnection()
	if error != nil {
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	rowsAffected, err := repository.Delete(id)
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}

// UpdatePassword updates a user password
func (s *userService) UpdatePassword(id uint64, updatePasswordDTO dtos.UpdatePassword) (rowsAffected int64, error *errors.Error) {
	db, error := database.OpenConnection()
	if error != nil {
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	currentPassword, err := repository.GetPassword(id)
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

	rowsAffected, err = repository.UpdatePassword(id, string(hash))
	if err != nil {
		return 0, errors.NewError(err.Error(), http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return 0, errors.NewError("no rows affected", http.StatusBadRequest)
	}

	return
}
