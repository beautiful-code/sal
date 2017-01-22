package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/user/controllers"
)

// InitRoutes registers all routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = setUserRoutes(router)

	return router
}

func setUserRoutes(router *mux.Router) *mux.Router {

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
