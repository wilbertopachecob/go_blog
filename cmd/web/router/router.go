package router

import (
	"wilbertopachecob/go_blog/cmd/web/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutes(r)
}
