package dtos

type AddressDto struct {
	ID         uint64 `json:"id,omitempty"`
	UserID     uint64 `json:"user_id,omitempty"`
	Complement string `json:"complement,omitempty"`
	Number     uint64 `json:"number,omitempty"`
	Cep        string `json:"cep,omitempty"`
	CityID     uint64 `json:"city_id,omitempty"`
}

type CreateAddressDto struct {
	UserID     uint64 `json:"user_id,omitempty"`
	Complement string `json:"complement,omitempty"`
	Number     uint64 `json:"number,omitempty"`
	Cep        string `json:"cep,omitempty"`
	CityID     uint64 `json:"city_id,omitempty"`
}
