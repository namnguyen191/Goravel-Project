package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/namnguyen191/goravel"
)

func initApplication() *application {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	// init goravel
	grv := &goravel.Goravel{}
	err = grv.New(path)
	if err != nil {
		log.Fatal(err)
	}

	grv.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: grv,
	}

	myHandlers := &handlers.Handlers{
		App: grv,
	}

	grv.InfoLog.Println("Debug is set to", grv.Debug)

	app := &application{
		App:        grv,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
