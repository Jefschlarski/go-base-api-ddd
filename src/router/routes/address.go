package routes

import (
	"api/src/controllers"
	"net/http"
)

var addressRoutes = []Route{
	{
		Uri:     "/address",
		Method:  http.MethodPost,
		Handler: controllers.CreateAddress,
		ReqAuth: true,
	},
	{
		Uri:    "/user/{user_id}/address",
		Method: http.MethodGet,
		// Handler: controllers.GetAddressByUserID,
		ReqAuth: true,
	},
}
