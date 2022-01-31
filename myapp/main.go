package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/namnguyen191/goravel"
)

type application struct {
	App        *goravel.Goravel
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	grv := initApplication()

	grv.App.ListenAndServe()
}
