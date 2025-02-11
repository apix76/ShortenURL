package UseCase

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

func ShortenURL(URL string) string {
	hash := md5.Sum([]byte(URL))

	encoded := base64.URLEncoding.EncodeToString(hash[:])
	encoded = strings.ReplaceAll(encoded, "+", "_")
	encoded = strings.ReplaceAll(encoded, "/", "_")

	shortURL := encoded[:10]
	return shortURL
}
