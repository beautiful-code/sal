package main

import (
	//"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/user/app"
	"github.com/beautiful-code/sal/services/user/models"
	"github.com/beautiful-code/sal/services/user/routers"
)

// Entry point of the program
func main() {
	app.InitData()

	common.InitKeys()
	utils.SetLogLevel(utils.Level(app.Data.Config.LogLevel))

	// Run DB Migrations
	app.Data.DB.AutoMigrate(&model.User{})

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
