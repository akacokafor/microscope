.PHONY: build dev compile-assets

build: compile-assets
	go build -v -ldflags="-s -w -X github.com/akacokafor/microscope/cmd.AppEnv=prod" -o bin/app main.go

compile-assets:
	npm run production

dev: 
	npm run watch & go run main.go