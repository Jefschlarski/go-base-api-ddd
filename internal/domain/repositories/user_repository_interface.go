package repositories

import "taskmanager/internal/domain/entities"

type UserRepositoryInterface interface {
	Create(user *entities.User) (lastInsertID uint64, err error)
	GetAll() (userList []entities.User, err error)
	Get(id uint64) (user entities.User, err error)
	Update(id uint64, user entities.User) (rowsAffected int64, err error)
	Delete(id uint64) (rowsAffected int64, err error)
	UpdatePassword(id uint64, userPassword string) (rowsAffected int64, err error)
	GetPassword(id uint64) (password string, err error)
	GetByEmail(email string) (user entities.User, err error)
}
