package url

import (
	"url-shortner/db"

	"github.com/go-redis/redis/v8"
)

func checkKeyExist(key string) bool {
	rdb := db.GetDbInstance()

	_, err := rdb.GetValue(key)

	if err == redis.Nil {
		return false
	}
	return true
}

func setUrl(key string, value string) error {
	rdb := db.GetDbInstance()
	err := rdb.SetValue(key, value)

	return err
}

func getUrl(key string) (string, error) {
	rdb := db.GetDbInstance()
	res, err := rdb.GetValue(key)

	return res, err
}
