package repositories

import "taskmanager/internal/domain/entities"

type UserRepositoryInterface interface {
	Create(user *entities.User) (uint64, error)
	GetAll() ([]entities.User, error)
	Get(id uint64) (entities.User, error)
	Update(id uint64, user entities.User) (rowsAffected int64, err error)
	Delete(id uint64) (rowsAffected int64, err error)
	UpdatePassword(id uint64, userPassword string) (rowsAffected int64, err error)
	GetPassword(id uint64) (string, error)
	GetByEmail(email string) (entities.User, error)
}
