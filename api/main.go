package main

import (
	"log"
	"url-shortner/config"
	"url-shortner/db"
	"url-shortner/webserver"
)

func main() {

	//initialise config
	appConfig := config.InitConfig()

	//initialise db
	err := db.InitDb(appConfig.DbConfig)

	if err != nil {
		log.Fatal(err)
	}

	//init server
	webServer := webserver.InitWebServer(&appConfig.WebConfig)

	//start server
	webServer.Serve()
}
