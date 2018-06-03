package utils

import (
	"errors"

	"github.com/go-redis/redis"
	"github.com/my-beego-todo/models"
)

const COUNTERKEY = "shorturl::counter"
const INITCOUNT = 10000

func InitCounter() {
	var cache = models.GetRedis()
	_, err := cache.Get(COUNTERKEY).Result()
	if err == redis.Nil {
		cache.Set(COUNTERKEY, INITCOUNT, 0)
	} else if err != nil {
		errors.New("Redis init count failed")
	}
}

func GetNextSequence() int {
	var cache = models.GetRedis()

}

func RunSequence() {
	for {
		select {}
	}
}
