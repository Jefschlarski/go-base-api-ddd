package entities

import (
	"net/http"
	"strings"
	"taskmanager/internal/common/errors"
	"taskmanager/internal/common/security"
	"taskmanager/internal/domain/valueobjects"
)

// User struct represents a user in the database
type User struct {
	ID       uint64                `json:"id,omitempty"`
	Name     string                `json:"name,omitempty"`
	Cpf      valueobjects.Cpf      `json:"cpf,omitempty"`
	Type     valueobjects.UserType `json:"type,omitempty"`
	Phone    valueobjects.Phone    `json:"phone,omitempty"`
	Password string                `json:"password,omitempty"`
	Email    valueobjects.Email    `json:"email,omitempty"`
}

// Prepare prepares the user for further processing.
//
// It takes a boolean parameter `isCreate` which indicates whether the user is being created or not.
// It returns an error if there is any validation error or nil if the user is prepared successfully.
func (user *User) Prepare(isCreate bool) *errors.Error {
	user.formater()
	if err := user.validate(isCreate); err != "" {
		return errors.NewError(err, http.StatusBadRequest)
	}

	if isCreate {

		hash, err := security.Hash(user.Password)
		if err != nil {
			return errors.NewError(err.Error(), http.StatusInternalServerError)
		}

		user.Password = string(hash)
	}

	return nil
}

// validate checks if the user fields are empty and returns an error if any field is empty.
//
// The function takes a boolean parameter `isCreate` which indicates whether the user is being created or not.
// It returns an error string if any required field is empty, otherwise it returns nil.
func (user *User) validate(isCreate bool) string {
	if user.Name == "" {
		return "name is required"
	}
	if user.Cpf == "" {
		return "cpf is required"
	}
	if err := user.Cpf.Validate(); err != nil {
		return err.Error()
	}
	if user.Phone == "" {
		return "phone is required"
	}
	if err := user.Phone.Validate(); err != nil {
		return err.Error()
	}
	if isCreate && user.Password == "" {
		return "password is required"
	}
	if user.Email == "" {
		return "email is required"
	}
	if err := user.Email.Validate(); err != nil {
		return err.Error()
	}

	return ""
}

// formater remove empty spaces in user fields
func (user *User) formater() {
	user.Name = strings.TrimSpace(user.Name)
	user.Cpf = user.Cpf.Formater()
	user.Phone = user.Phone.Formater()
	user.Email = user.Email.Formater()
}
