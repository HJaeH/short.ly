package models

import (
	"encoding/json"
	"fmt"

	"github.com/short.ly/db/redis"
	"github.com/short.ly/utils/base62"
	"github.com/short.ly/utils/url_organizer"
)

const URL_KEY_PRE = "shorturl::url"
const COUNTER_KEY = "shorturl::counter"

func SetUrl(shortUrl string, originalUrl string) error {
	db := redis.GetRedis()
	_, err := db.Set(shortUrl, originalUrl, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetUrl(url string) (*URL, error) {
	key := createURLKey(url)
	db := redis.GetRedis()

	val, err := db.Get(key).Result()
	var newURL = new(URL)
	if err != nil {
		return &URL{}, err
	}

	if err := json.Unmarshal([]byte(val), &newURL); err != nil {
		return &URL{}, err
	}

	return newURL, nil
}

//func GetCount() (int64, error) {
//
//	cache := redis.GetRedis()
//	var count = new(int64)
//	res, err := cache.Get(COUNTER_KEY).Result()
//	if err != nil {
//		cache.Incr(COUNTER_KEY)
//	}
//	if err := json.Unmarshal([]byte(res), &count); err != nil {
//		return 0, err
//	}
//
//	return *count, nil
//}

func AddURL(url string) error {

	uniqueURL := url_organizer.GetUniqueURL(url)

	index, err := redis.GetIndex()
	if err != nil {
		return err
	}
	shortURL := base62.EncodeBase62(int64(index))
	fmt.Println("shorturl test", shortURL)
	urlAlreadyExist, err := GetUrl(shortURL)
	if err == redis.Nil {
		SetUrl(shortURL, uniqueURL)
	}
	_ = urlAlreadyExist
	return fmt.Errorf("unknown error")
}

/// internal functions
func createURLKey(organizedURL string) string {
	return URL_KEY_PRE + "::" + organizedURL
}
