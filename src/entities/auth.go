package entities

// Auth struct represents an authentication request struct
type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
