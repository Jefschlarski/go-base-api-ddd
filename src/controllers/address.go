package controllers

import (
	"net/http"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {

	// var address entities.Address

	// if err := request.ProcessBody(r, &address); err != nil {
	// 	responses.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// if err := security.VerifyId(address.UserID, r); err != nil {
	// 	responses.Error(w, http.StatusForbidden, err)
	// 	return
	// }

	// db, err := database.OpenConnection()
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// defer db.Close()

	// // repository := repositories.NewAddressRepository(db)

	// // address.ID, err = repository.Create(address)
	// // if err != nil {
	// // 	responses.Error(w, http.StatusInternalServerError, err)
	// // 	return
	// // }

	// responses.Json(w, http.StatusCreated, address)
}
