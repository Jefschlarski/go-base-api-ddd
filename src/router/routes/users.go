package routes

import (
	"api/src/handlers"
	"net/http"
)

var userRoutes = []Routes{

	{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: handlers.CreateUser,
		ReqAuth: false,
	},
	{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: handlers.GetUsers,
		ReqAuth: false,
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: handlers.GetUser,
		ReqAuth: false,
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: handlers.UpdateUser,
		ReqAuth: false,
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: handlers.DeleteUser,
		ReqAuth: false,
	},
}
