package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	WebConfig WebConfig
	DbConfig  DbConfig
}

type WebConfig struct {
	Port string `default:"8080"`
}

type DbConfig struct {
	Port    string `default:"6379"`
	Address string `default:"localhost:6379"`
}

var envFile = "local.env"

func InitConfig() *AppConfig {

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("error at loading .env: ", err)
	}

	webPort := os.Getenv("WEB_PORT")
	dbPort := os.Getenv("REDIS_PORT")
	dbAdress := os.Getenv("REDIS_ADRESS")

	webConfig := WebConfig{Port: webPort}
	dbConfig := DbConfig{
		Port:    dbPort,
		Address: dbAdress,
	}

	return &AppConfig{
		WebConfig: webConfig,
		DbConfig:  dbConfig,
	}
}
