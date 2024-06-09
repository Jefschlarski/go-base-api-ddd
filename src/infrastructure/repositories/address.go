package repositories

import (
	repositoriesInterfaces "api/src/domain/repositories"
	"api/src/infrastructure/database"
	"api/src/interface/api/dtos"
)

// address struct represents a address repository
type addressRepository struct{}

// NewAddressRepository create a new address repository
func NewAddressRepository() repositoriesInterfaces.AddressRepository {
	return &addressRepository{}
}

// Create insert a new address in the database
func (u addressRepository) Create(address dtos.CreateAddressDto) (lastInsertID uint64, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	statement, err := db.Prepare(`insert into "address" (user_id, complement, number, cep, city_id) values ($1, $2, $3, $4, $5) returning id`)
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
func (u addressRepository) GetByUserID(userID uint64) (addressess []dtos.AddressDto, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(`select id, user_id, complement, number, cep, city_id from "address" where user_id = $1`, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address dtos.AddressDto
		if err = rows.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID); err != nil {
			return
		}
		addressess = append(addressess, address)
	}

	return
}

// UpdateAddressesByID updates a user address by address id
func (u addressRepository) Update(address dtos.AddressDto) (rowsAffected int64, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	statement, err := db.Prepare(`update "address" set complement = $1, number = $2, cep = $3, city_id = $4 where id = $5`)
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
func (u addressRepository) Get(id uint64) (address dtos.AddressDto, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	row, err := db.Query(`select id, user_id, complement, number, cep, city_id from "address" where id = $1`, id)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		if err = row.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID); err != nil {
			return
		}
	}

	return
}

// GetAddresses get all addresses
func (u addressRepository) GetAll() (addressess []dtos.AddressDto, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(`select id, user_id, complement, number, cep, city_id from "address"`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var address dtos.AddressDto
		if err = rows.Scan(&address.ID, &address.UserID, &address.Complement, &address.Number, &address.Cep, &address.CityID); err != nil {
			return
		}
		addressess = append(addressess, address)
	}

	return
}

// DeleteAddressByID delete a address by id
func (u addressRepository) Delete(id uint64) (rowsAffected int64, err error) {

	db, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer db.Close()

	statement, err := db.Prepare(`delete from "address" where id = $1`)
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
