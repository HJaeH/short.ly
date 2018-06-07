package models

import (
	"github.com/HJaeH/short.ly/db/redis"
	"github.com/HJaeH/short.ly/utils/base62"
	"github.com/HJaeH/short.ly/utils/url_organizer"
)

var UrlList *UrlManager

func SetUrl(shortUrl string, originalUrl string) error {
	db := redis.GetRedis()
	//bi-direction
	pipe := db.Pipeline()
	pipe.Set(originalUrl, shortUrl, 0)
	pipe.Set(shortUrl, originalUrl, 0)
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func GetOriginalUrl(shortUrl string) (string, error) {
	key := shortUrl
	db := redis.GetRedis()

	originalUrl, err := db.Get(key).Result()
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func AddURL(url string) (string, error) {
	db := redis.GetRedis()
	uniqueURL := url_organizer.GetUniqueURL(url)
	existShortUrl, err := db.Get(uniqueURL).Result()
	var shortURL string
	if err == redis.Nil {
		index, err := redis.GetIndex()
		if err != nil {
			return "", err
		}
		shortURL = base62.EncodeBase62(int64(index))

		if err := SetUrl(shortURL, uniqueURL); err != nil {
			return "", err
		}
		return shortURL, nil
	}

	return existShortUrl, nil
}
