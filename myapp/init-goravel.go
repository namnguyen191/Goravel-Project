package main

import (
	"log"
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

	grv.InfoLog.Println("Debug is set to", grv.Debug)

	app := &application{
		App: grv,
	}

	return app
}
