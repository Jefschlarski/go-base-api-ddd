package dtos

// CityDto represents an CityDto struct

type CityDto struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	StateID uint64 `json:"state_id,omitempty"`
}
