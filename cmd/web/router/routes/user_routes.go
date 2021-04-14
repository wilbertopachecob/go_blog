package routes

import (
	"net/http"
	"wilbertopachecob/go_blog/cmd/web/controllers"
)

var UserRoutes = []Route{
	{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	{
		Uri:     "/user",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
}
