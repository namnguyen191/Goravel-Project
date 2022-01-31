BINARY_NAME=goravelApp.exe

## build: builds all binaries
build:
	@go mod vendor
	@echo "building Goravel"
	@go build -o tmp/${BINARY_NAME} .
	@echo "Goravel built!""

run: build
	@echo Staring Goravel...
	@.\tmp\${BINARY_NAME} &
	@echo Goravel started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
	
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Celeritas

restart: stop start