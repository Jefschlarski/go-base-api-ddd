package interfaces

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/domain/entities"
)

type CreateUser interface {
	Execute(user *entities.User) (uint64, *errors.Error)
}

type GetAllUsers interface {
	Execute() ([]entities.User, *errors.Error)
}

type GetUser interface {
	Execute(id uint64) (entities.User, *errors.Error)
}

type UpdateUser interface {
	Execute(id uint64, user entities.User) (rowsAffected int64, err *errors.Error)
}

type DeleteUser interface {
	Execute(id uint64) (rowsAffected int64, err *errors.Error)
}
type UpdateUserPassword interface {
	Execute(id uint64, updatePasswordDTO dtos.UpdatePassword) (rowsAffected int64, error *errors.Error)
}

type GetUserPassword interface {
	GetPassword(id uint64) (string, err *errors.Error)
}
