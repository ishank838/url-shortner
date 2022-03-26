package url

import (
	"fmt"
	"url-shortner/model"
	keygenerator "url-shortner/pkg/keyGenerator"

	"github.com/go-redis/redis/v8"
)

func ShortenUrl(url model.ShortenRequest) (*model.ShortenResponse, error) {
	key := keygenerator.GenerateKey(url.Url)

	if checkKeyExist(key) {
		return &model.ShortenResponse{
			Url:        url.Url,
			ShortenUrl: buildShortenUrl(key),
		}, nil
	}

	err := setUrl(key, url.Url)

	if err != nil {
		return nil, err
	}

	return &model.ShortenResponse{
		Url:        url.Url,
		ShortenUrl: buildShortenUrl(key),
	}, nil
}

func ResolveUrl(hash string) (string, error) {

	resp, err := getUrl(hash)

	if err == redis.Nil {
		return "", fmt.Errorf("invalid url")
	}

	if err != nil {
		return "", fmt.Errorf("error at resolve url")
	}

	return resp, nil
}

func buildShortenUrl(hash string) string {
	return "localhost:8080/" + hash
}
