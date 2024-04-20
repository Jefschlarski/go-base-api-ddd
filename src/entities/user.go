package entities

import (
	"api/src/common/security"
	"api/src/common/validate"
	"errors"
	"strings"
	"time"
)

// User struct represents a user in the database
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Type      uint      `json:"type,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Prepare prepares the user for further processing.
//
// It takes a boolean parameter `isCreate` which indicates whether the user is being created or not.
// It returns an error if there is any validation error or nil if the user is prepared successfully.
func (user *User) Prepare(isCreate bool) error {
	user.formater()
	if err := user.validate(isCreate); err != nil {
		return err
	}

	if isCreate {

		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hash)
	}

	return nil
}

// validate checks if the user fields are empty and returns an error if any field is empty.
//
// The function takes a boolean parameter `isCreate` which indicates whether the user is being created or not.
// It returns an error if any required field is empty, otherwise it returns nil.
func (user *User) validate(isCreate bool) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Cpf == "" {
		return errors.New("cpf is required")
	}
	if user.Type == 0 {
		return errors.New("type is required")
	}
	if user.Phone == "" {
		return errors.New("phone is required")
	}
	if isCreate && user.Password == "" {
		return errors.New("password is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	if err := validate.Email(user.Email); err != nil {
		return err
	}

	return nil
}

// formater remove empty spaces in user fields
func (user *User) formater() {
	user.Name = strings.TrimSpace(user.Name)
	user.Cpf = strings.TrimSpace(user.Cpf)
	user.Phone = strings.TrimSpace(user.Phone)
	user.Email = strings.TrimSpace(user.Email)
}
