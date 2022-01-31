package middleware

import (
	"myapp/data"

	"github.com/namnguyen191/goravel"
)

type Middleware struct {
	App    *goravel.Goravel
	Models data.Models
}
