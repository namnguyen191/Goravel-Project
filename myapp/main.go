package main

import (
	"github.com/namnguyen191/goravel"
)

type application struct {
	App *goravel.Goravel
}

func main() {
	grv := initApplication()

	grv.App.ListenAndServe()
}
