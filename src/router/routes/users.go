package routes

import (
	"api/src/handlers"
	"net/http"
)

var userRoutes = []Route{

	{
		Uri:     "/user",
		Method:  http.MethodPost,
		Handler: handlers.CreateUser,
		ReqAuth: false,
	},
	{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: handlers.GetUsers,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: handlers.GetUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: handlers.UpdateUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: handlers.DeleteUser,
		ReqAuth: true,
	},
	{
		Uri:     "/user/{id}/update-password",
		Method:  http.MethodPost,
		Handler: handlers.UpdateUserPassword,
		ReqAuth: true,
	},
}
