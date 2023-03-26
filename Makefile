# Makefile

build: # * build server
	go build -o ./.bin/app ./cmd/api/main.go

build-linux: # * build server for Linux
	GOOS=linux GOARCH=amd64 go build -o ./.bin/app ./cmd/api/main.go

start: # * start server
	./.bin/app

dev: # * build and start server
	go build -o ./.bin/app ./cmd/api/main.go
	./.bin/app
