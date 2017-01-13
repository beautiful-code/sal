package routers

import (
	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/user/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	authRouter := mux.NewRouter()
	authRouter.HandleFunc("/user", controllers.GetUser).Methods("GET")

	router.PathPrefix("/user").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(authRouter),
	))

	return router
}
