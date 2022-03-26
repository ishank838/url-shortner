package redis

import (
	"fmt"
	"log"
	"url-shortner/config"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type DbRedis struct {
	rdbClient *redis.Client
}

var redisDb *DbRedis

func GetRedisInstance() *DbRedis {
	return redisDb
}

func InitRedisDb(dbConfig config.DbConfig) error {
	rdClient := redis.NewClient(&redis.Options{
		Addr: dbConfig.Address,
	})

	_, err := rdClient.Ping(context.TODO()).Result()

	if err != nil {
		return fmt.Errorf("error at ping redis: %v", err)
	}

	log.Println("connected redis at : ", dbConfig.Port)

	redisDb = &DbRedis{
		rdbClient: rdClient,
	}
	return nil
}
