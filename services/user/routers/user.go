package routers

import (
	"github.com/beautiful-code/sal/services/user/controllers"

	"github.com/gorilla/mux"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
