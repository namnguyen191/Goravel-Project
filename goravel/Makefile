## test: run all tests
test:
	@go test -v ./...

## cover: open coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover --html=coverage.out

## coverage: displays tests coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool Goravel and copies it to myapp
build_cli:
	@go build -o ../myapp/goravel.exe ./cmd/cli