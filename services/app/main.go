package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/app/models"
	"github.com/beautiful-code/sal/services/app/routers"
)

func main_init(configFilePath string) {
	common.InitConfig(configFilePath)
	common.SetLogLevel(common.Level(common.AppConfig.LogLevel))
	common.CreateDBSession()
}

// Entry point of the program
func main() {
	main_init("config.json")

	// TODO Move this logic to a better place
	dataStore := common.NewDataStore()
	dataStore.Session.AutoMigrate(&model.Application{})
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

	log.Println("Listening on http://localhost:8090 ...")

	// Running the HTTP Server
	server.ListenAndServe()
}
