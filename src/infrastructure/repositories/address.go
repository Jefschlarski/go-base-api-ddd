package repositories

import (
	"api/src/api/dtos"
	"api/src/domain/entities"
	"api/src/infrastructure/database"
)

type AddressRepositoryInterface interface {
	Create(dtos.CreateAddressDto) (uint64, error)
	GetByUserID(uint64) ([]entities.Address, error)
	Update(dtos.AddressDto) (int64, error)
	Get(uint64) (entities.Address, error)
	GetAll() ([]entities.Address, error)
	Delete(id uint64) (int64, error)
}

// address struct represents a address repository
type addressRepository struct {
	db database.DatabaseInterface
}

// NewAddressRepository create a new address repository
func NewAddressRepository(db database.DatabaseInterface) AddressRepositoryInterface {
	return &addressRepository{db}
}

// Create insert a new address in the database
func (a addressRepository) Create(address dtos.CreateAddressDto) (lastInsertID uint64, err error) {

	statement, err := a.db.Prepare(`insert into "address" (user_id, complement, number, cep, city_id) values ($1, $2, $3, $4, $5) returning id`)
	if err != nil {
		return
	}
	defer statement.Close()

	err = statement.QueryRow(address.UserID, address.Complement, address.Number, address.Cep, address.CityID).Scan(&lastInsertID)
	if err != nil {
		return
	}

	return
}

// GetAddressesByUserID get all addresses by user id
func (a addressRepository) GetByUserID(userID uint64) (addressess []entities.Address, err error) {

	rows, err := a.db.Query(`select id, user_id, complement, number, cep, city_id, created_at, updated_at from "address" where user_id = $1`, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address entities.Address
		if err = rows.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID, &address.CreatedAt, &address.UpdatedAt); err != nil {
			return
		}
		addressess = append(addressess, address)
	}

	return
}

// UpdateAddressesByID updates a user address by address id
func (a addressRepository) Update(address dtos.AddressDto) (rowsAffected int64, err error) {

	statement, err := a.db.Prepare(`update "address" set complement = $1, number = $2, cep = $3, city_id = $4 where id = $5`)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(address.Complement, address.Number, address.Cep, address.CityID, address.ID)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}

// GetAddressByID get a address by id
func (a addressRepository) Get(id uint64) (address entities.Address, err error) {

	row, err := a.db.Query(`select id, user_id, complement, number, cep, city_id, created_at, updated_at from "address" where id = $1`, id)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		if err = row.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID, &address.CreatedAt, &address.UpdatedAt); err != nil {
			return
		}
	}

	return
}

// GetAddresses get all addresses
func (a addressRepository) GetAll() (addressess []entities.Address, err error) {

	rows, err := a.db.Query(`select id, user_id, complement, number, cep, city_id, created_at, updated_at from "address"`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address entities.Address
		if err = rows.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID, &address.CreatedAt, &address.UpdatedAt); err != nil {
			return
		}
		addressess = append(addressess, address)
	}

	return
}

// DeleteAddressByID delete a address by id
func (a addressRepository) Delete(id uint64) (rowsAffected int64, err error) {

	statement, err := a.db.Prepare(`delete from "address" where id = $1`)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}
