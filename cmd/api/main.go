package main

import (
	"log"
	"static/server"

	"static/pkg/config"
)

func main() {
	conf := config.GetConfig()

	app := server.NewApp()

	if err := app.Run(conf.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
