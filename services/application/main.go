package main

import (
	"net/http"

	"github.com/rs/cors"
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

	app.Data.DB.AutoMigrate(&models.Application{})
	app.Data.DB.AutoMigrate(&models.Feedback{})

	// Get the mux router object
	router := routers.InitRoutes()

	// Handle CORS
	c := cors.New(cors.Options{
		AllowedOrigins: app.Data.Config.AllowedOrigins,
	})

	// Create a negroni instance
	n := negroni.Classic()
	n.Use(c)
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
