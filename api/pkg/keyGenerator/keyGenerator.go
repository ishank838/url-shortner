package keygenerator

import (
	"crypto/md5"
	"encoding/base64"
)

func GenerateKey(url string) string {
	hash := makeHash(url)

	key := base64.StdEncoding.EncodeToString(hash)

	return key
}

func makeHash(url string) []byte {
	hash := md5.Sum([]byte(url))
	return hash[:]
}
