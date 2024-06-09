package routes

import (
	"api/src/interface/api/controllers"
	"net/http"
)

var addressRoutes = []Route{
	{
		Uri:     "/user/{user_id}/address",
		Method:  http.MethodPost,
		Handler: controllers.CreateAddress,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{user_id}/address",
		Method:  http.MethodGet,
		Handler: controllers.GetAddressesByUserID,
		ReqAuth: true,
	},
	{
		Uri:     "/address/{address_id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateAddressesByID,
		ReqAuth: true,
	},
	{
		Uri:     "/address/{address_id}",
		Method:  http.MethodGet,
		Handler: controllers.GetAddressById,
		ReqAuth: true,
	},
	{
		Uri:     "/address",
		Method:  http.MethodGet,
		Handler: controllers.GetAddresses,
		ReqAuth: true,
	},
	{
		Uri:     "/address/{address_id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteAddressesByID,
		ReqAuth: true,
	},
}
