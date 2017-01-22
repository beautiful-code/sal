package main

import (
	"net/http"

	"github.com/urfave/negroni"

	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/application/app"
	"github.com/beautiful-code/sal/services/application/models"
	"github.com/beautiful-code/sal/services/application/routers"
)

// Entry point of the program
func main() {
	app.InitData()

	utils.SetLogLevel(utils.Level(app.Data.Config.LogLevel))

	app.Data.DB.AutoMigrate(&model.Application{})

	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	// Create the Server
	server := &http.Server{
		Addr:    app.Data.Config.Server,
		Handler: n,
	}

	utils.Info.Printf("Listening on http://%s ...", app.Data.Config.Server)
	// Running the HTTP Server
	server.ListenAndServe()
}
