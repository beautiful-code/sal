package routers

import (
	"github.com/beautiful-code/sal/services/app/controllers"

	"github.com/gorilla/mux"
)

func SetApplicationRoutes(router *mux.Router) *mux.Router {

	//router.HandleFunc("/applications", controllers.List).Methods("POST")
	router.HandleFunc("/applications/create", controllers.Create).Methods("POST")

	return router
}
