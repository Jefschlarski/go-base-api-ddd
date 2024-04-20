package repositories

import (
	"api/src/entities"
	"database/sql"
)

// users struct represents a user repository
type users struct {
	db *sql.DB
}

// NewUsersRepository create a new user repository
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insert a new user in the database
func (u users) Create(user entities.User) (uint64, error) {

	statement, err := u.db.Prepare("insert into users (name, cpf, type, phone, password, email) values ($1, $2, $3, $4, $5, $6) returning id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertID uint64
	err = statement.QueryRow(user.Name, user.Cpf, user.Type, user.Phone, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
