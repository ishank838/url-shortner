package url

import (
	"url-shortner/model"
	keygenerator "url-shortner/pkg/keyGenerator"
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

func buildShortenUrl(hash string) string {
	return "localhost:8080/" + hash
}
