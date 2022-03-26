package db

import (
	"fmt"
	"log"
	"url-shortner/config"
	"url-shortner/db/redis"
)

func InitDb(dbConfig config.DbConfig) error {
	log.Println("connecting to redis at: ", dbConfig.Port)

	err := redis.InitRedisDb(dbConfig)

	if err != nil {
		return fmt.Errorf("error at initDb: %v", err)
	}

	return nil
}
