package repositories

import (
	"database/sql"
	"fmt"
	"taskmanager/internal/domain/entities"
	"taskmanager/internal/domain/repositories"
)

// user struct represents a user repository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository create a new user repository
func NewUserRepository(db *sql.DB) repositories.UserRepositoryInterface {
	return &userRepository{db}
}

// Create insert a new user in the database
func (u userRepository) Create(user *entities.User) (lastInsertID uint64, err error) {

	statement, err := u.db.Prepare(`insert into "user" (name, cpf, type, phone, password, email) values ($1, $2, $3, $4, $5, $6) returning id`)
	if err != nil {
		return
	}
	defer statement.Close()

	err = statement.QueryRow(user.Name, user.Cpf, user.Type, user.Phone, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		return
	}

	return
}

// Get get a user by ID
func (u userRepository) Get(id uint64) (user entities.User, err error) {

	rows, err := u.db.Query(`select name, cpf, type, phone, email from "user" where id = $1`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&user.Name, &user.Cpf, &user.Type, &user.Phone, &user.Email); err != nil {
			return
		}
	} else {
		err = fmt.Errorf("there is no user with id %d", id)
		return
	}

	return
}

// GetAll get all users
func (u userRepository) GetAll() (userList []entities.User, err error) {

	rows, err := u.db.Query(`select id, name, cpf, type, phone, email from "user"`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Cpf, &user.Type, &user.Phone, &user.Email); err != nil {
			return
		}
		userList = append(userList, user)
	}

	return
}

// Update update a user
func (u userRepository) Update(id uint64, user entities.User) (rowsAffected int64, err error) {

	statement, err := u.db.Prepare(`update "user" set name = $1, cpf = $2, type = $3, phone = $4, email = $5 where id = $6`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Cpf, user.Type, user.Phone, user.Email, id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

// Delete delete a user
func (u userRepository) Delete(id uint64) (rowsAffected int64, err error) {

	statement, err := u.db.Prepare(`delete from "user" where id = $1`)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

// GetByEmail get a user name, email and password by email
func (u userRepository) GetByEmail(email string) (user entities.User, err error) {
	rows, err := u.db.Query(`select id, name, email, password from "user" where email = $1`, email)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return
		}
	} else {
		err = fmt.Errorf("there is no user with email %s", email)
		return
	}

	return
}

// getPassword get a user password by id
func (u userRepository) GetPassword(id uint64) (password string, err error) {
	rows, err := u.db.Query(`select password from "user" where id = $1`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&password); err != nil {
			return
		}
	}

	return
}

// UpdatePassword update a user password
func (u userRepository) UpdatePassword(id uint64, password string) (rowsAffected int64, err error) {

	statement, err := u.db.Prepare(`update "user" set password = $1 where id = $2`)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(password, id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}
