package entities

// address struct represents an address Aggregate
type address struct {
	ID         uint64 `json:"id,omitempty"`
	User       User   `json:"user_id,omitempty"`
	Complement string `json:"complement,omitempty"`
	Number     uint64 `json:"number,omitempty"`
	Cep        string `json:"cep,omitempty"`
	City       City   `json:"city,omitempty"`
}

// NewAddress
func NewAddress(user User, complement string, number uint64, cep string, city City) address {
	return address{
		User:       user,
		Complement: complement,
		Number:     number,
		Cep:        cep,
		City:       city,
	}
}
