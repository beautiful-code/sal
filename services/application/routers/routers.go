package routers

import (
	"github.com/gorilla/mux"

	//"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/application/controllers"
)

func SetApplicationRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/applications", controllers.List).Methods("POST")
	router.HandleFunc("/applications/create", controllers.Create).Methods("POST")
	router.HandleFunc("/feedbacks/create", controllers.CreateFeedback).Methods("POST")
	router.HandleFunc("/feedbacks", controllers.ListFeedbacks).Methods("POST")

	return router
}

// InitRoutes registers all routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetApplicationRoutes(router)

	return router
}
