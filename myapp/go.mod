module myapp

go 1.17

replace github.com/namnguyen191/goravel => ../goravel

require github.com/namnguyen191/goravel v0.0.0-00010101000000-000000000000

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
