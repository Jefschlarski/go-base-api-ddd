package interfaces

import (
	"taskmanager/internal/common/errors"
	"taskmanager/internal/domain/entities"
)

type GetAllStates interface {
	Execute() ([]entities.State, *errors.Error)
}
