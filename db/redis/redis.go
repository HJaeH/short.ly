package redis

import (
	"errors"

	"sync"

	"math"
	"math/rand"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	database "github.com/HJaeH/short.ly/db"
	"github.com/HJaeH/short.ly/utils/base62"
	"github.com/HJaeH/short.ly/utils/result_code"
)

const Nil = redis.Nil

const StartCharacterLength = 3


func newRedisClient() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		errors.New("Redis connection failed")
		return nil, err
	}
	return client, nil

}

type RedisClient struct {
	*redis.Client
}

var redisClient *redis.Client
var gOnceCache sync.Once

func GetRedis() *redis.Client {
	gOnceCache.Do(func() {
		redisClient, _ = newRedisClient()
		//init numbers
		return
	})
	return redisClient
}

func ExpendCount() error {
	//exist, err := isExist(database.COUNT_KEY)
	//if err != nil {
	//	return err
	//}

	var startCount int
	var newMaxCount int
	db := GetRedis()
	maxCountResult, err := db.Get(database.MAX_COUNT_KEY).Result()
	var first bool = false
	if err == redis.Nil {
		first = true
		startCount = int(math.Pow(float64(base62.Base), StartCharacterLength-1))
		newMaxCount = int(math.Pow(float64(base62.Base), float64(StartCharacterLength)))
	} else if err != nil {
		return err
	} else {

		currentMaxCount, err := strconv.Atoi(maxCountResult)
		if err != nil {
			return err
		}
		startCount = currentMaxCount
		newMaxCount = currentMaxCount + int(math.Pow(float64(base62.Base), float64(StartCharacterLength)))
	}


	pipe := db.Pipeline()
	for i := startCount; i < newMaxCount; i++ {
		pipe.LPush(database.COUNT_KEY, i)

	}
	if first {
		newMaxCount = newMaxCount - startCount
	}

	pipe.Set(database.MAX_COUNT_KEY, newMaxCount, 0).Result()
	_, err = pipe.Exec()
	if err != nil {
		return err
	}

	return nil
}

func GetIndex() (uint, error) {
	db := GetRedis()
	indexCandidatesCount, err := db.LLen(database.COUNT_KEY).Result()
	if err != nil {
		return 0, err
	}

	if indexCandidatesCount < 10000 {
		ExpendCount()
	}

	var randomInt = rand.Intn(int(indexCandidatesCount))


	randomIndex, err := db.LIndex(database.COUNT_KEY, int64(randomInt)).Result()
	if err != nil {
		return 0, err
	}

	result, err := db.LRem(database.COUNT_KEY, 0, randomIndex).Result()
	if err != nil {
		return 0, err
	}
	if result != 1 {
		return 0, errors.New(resultcode.ResultCodeMap[resultcode.ErrorDBKeyNotExist])
	}

	res, err := strconv.Atoi(randomIndex)
	if err != nil {
		return 0, err
	}

	return uint(res), nil

}

func isExist(key string) (bool, error) {
	db := GetRedis()
	result, err := db.Exists(key).Result()
	if err != nil {
		beego.Error(err, resultcode.ResultCodeMap[resultcode.ErrorDBConnection])
		return false, err
	}
	if result != 1 {
		return false, nil
	}

	return true, nil
}
