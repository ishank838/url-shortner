package main

import (
	"log"
	"url-shortner/config"
	"url-shortner/db"
)

func main() {
	appConfig := config.InitConfig()

	err := db.InitDb(appConfig.DbConfig)

	if err != nil {
		log.Fatal(err)
	}
}
