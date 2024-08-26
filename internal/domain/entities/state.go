package entities

// State struct represents a State in the database
type State struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Uf   string `json:"uf,omitempty"`
}
