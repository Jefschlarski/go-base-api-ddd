package controllers

import (
	"api/src/api/dtos"
	"api/src/application/common/errors"
	"api/src/application/common/request"
	"api/src/application/common/responses"
	"api/src/application/common/security"
	addressServices "api/src/application/services/address"
	cityServices "api/src/application/services/city"
	"api/src/infrastructure/database"
	"api/src/infrastructure/repositories"
	"net/http"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {

	var address dtos.CreateAddressDto
	err := request.ProcessBody(r, &address)
	if err != nil {
		responses.Error(w, err)
		return
	}

	userID, error := security.ExtractUserID(r)
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	address.UserID = userID

	db, error := database.NewDatabase()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	AddressRepository := repositories.NewAddressRepository(db)
	CityRepository := repositories.NewCityRepository(db)

	service := addressServices.NewCreateAddress(AddressRepository, CityRepository)
	err = service.Create(&address)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusCreated, address)
}

func GetAddressesByUserID(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r, "user_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	if err := security.VerifyId(userID, r); err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	repo := repositories.NewAddressRepository(db)
	cityRepo := repositories.NewCityRepository(db)

	service := addressServices.NewGetAddressByUserId(repo, cityRepo)
	addresses, err := service.GetByUserID(userID)
	if err != nil {
		responses.Error(w, err)
		return
	}
	responses.Json(w, http.StatusOK, addresses)
}

func UpdateAddressesByID(w http.ResponseWriter, r *http.Request) {

	var address dtos.AddressDto

	err := request.ProcessBody(r, &address)
	if err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	cityRepositories := repositories.NewCityRepository(db)

	cityService := cityServices.NewGetCity(cityRepositories)
	_, err = cityService.Get(address.CityID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	address.ID, err = request.GetId(r, "address_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	if err := security.VerifyId(address.UserID, r); err != nil {
		responses.Error(w, err)
		return
	}

	repo := repositories.NewAddressRepository(db)

	service := addressServices.NewUpdateAddress(repo)

	rowsAffected, err := service.Update(&address)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, rowsAffected)
}

func GetAddressById(w http.ResponseWriter, r *http.Request) {

	addressID, err := request.GetId(r, "address_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	repo := repositories.NewAddressRepository(db)

	service := addressServices.NewGetAddress(repo)
	address, err := service.Get(addressID)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, address)
}

func GetAddresses(w http.ResponseWriter, r *http.Request) {

	db, error := database.NewDatabase()
	if error != nil {
		err := errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	repo := repositories.NewAddressRepository(db)

	service := addressServices.NewGetAllAddresses(repo)

	addresses, err := service.GetAll()
	if err != nil {
		responses.Error(w, err)
		return
	}
	responses.Json(w, http.StatusOK, addresses)
}

func DeleteAddressesByID(w http.ResponseWriter, r *http.Request) {

	addressId, err := request.GetId(r, "address_id")
	if err != nil {
		responses.Error(w, err)
		return
	}

	db, error := database.NewDatabase()
	if error != nil {
		err = errors.NewError(error.Error(), http.StatusInternalServerError)
		responses.Error(w, err)
		return
	}
	defer db.Close()

	repo := repositories.NewAddressRepository(db)

	service := addressServices.NewDeleteAddress(repo)

	rowsAffected, err := service.Delete(addressId)
	if err != nil {
		responses.Error(w, err)
		return
	}

	responses.Json(w, http.StatusOK, rowsAffected)
}
