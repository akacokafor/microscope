.PHONY: build dev compile-assets

build: compile-assets
	go mod tidy
	go mod download
	go build -v -ldflags="-s -w -X github.com/akacokafor/microscope/cmd.AppEnv=prod" -o bin/app main.go

compile-assets:
	npm install && npm run production

dev: 
	npm install && npm run watch & go run main.go