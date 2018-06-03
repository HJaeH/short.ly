package models

import (
	"errors"
	"fmt"

	"sync"

	"github.com/go-redis/redis"
)

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		errors.New("Redis connection failed")
		return nil, err
	}
	fmt.Println(pong, err)
	redis.Nil
	return client, nil

}

var redisClient *redis.Client
var gOnceCache sync.Once

func GetRedis() *redis.Client {
	gOnceCache.Do(func() {
		redisClient, _ = newRedisClient()
		return
	})
	return redisClient
}
