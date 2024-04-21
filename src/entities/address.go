package entities

import "time"

// Address struct represents an address in the database
type Address struct {
	ID         uint64    `json:"id,omitempty"`
	UserID     uint64    `json:"user_id,omitempty"`
	Complement string    `json:"complement,omitempty"`
	Number     uint64    `json:"number,omitempty"`
	Cep        string    `json:"cep,omitempty"`
	City       uint64    `json:"city,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
