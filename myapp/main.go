package main

import (
	"myapp/handlers"

	"github.com/namnguyen191/goravel"
)

type application struct {
	App      *goravel.Goravel
	Handlers *handlers.Handlers
}

func main() {
	grv := initApplication()

	grv.App.ListenAndServe()
}
