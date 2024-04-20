package entities

import "time"

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
