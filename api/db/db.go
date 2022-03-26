package db

import (
	"fmt"
	"log"
	"url-shortner/config"
	"url-shortner/db/redisdb"
)

func InitDb(dbConfig config.DbConfig) error {
	log.Println("connecting to redis at: ", dbConfig.Port)

	err := redisdb.InitRedisDb(dbConfig)

	if err != nil {
		return fmt.Errorf("error at initDb: %v", err)
	}

	return nil
}

func GetDbInstance() *redisdb.DbRedis {
	return redisdb.GetRedisInstance()
}
