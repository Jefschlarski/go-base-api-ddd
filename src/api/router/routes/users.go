package routes

import (
	"api/src/api/controllers"
	"net/http"
)

var userRoutes = []Route{

	{
		Uri:     "/user",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
		ReqAuth: false,
	},
	{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}/update-password",
		Method:  http.MethodPost,
		Handler: controllers.UpdateUserPassword,
		ReqAuth: true,
	},
}
