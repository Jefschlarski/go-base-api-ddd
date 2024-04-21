package entities

// City struct represents a city in the database
type City struct {
	ID      uint64 `json:"id,omitempty"`
	StateID uint64 `json:"state_id,omitempty"`
	Name    string `json:"name,omitempty"`
	// CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt time.Time `json:"updated_at,omitempty"`
	State State
}
