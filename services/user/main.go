package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/user/models"
	"github.com/beautiful-code/sal/services/user/routers"
)

// Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp("config.json")

	// TODO Move this logic to a better place
	dataStore := common.NewDataStore()
	dataStore.Session.AutoMigrate(&model.User{})
	dataStore.Close()

	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	// Create the Server
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Listening on http://localhost:8080 ...")

	// Running the HTTP Server
	server.ListenAndServe()
}
