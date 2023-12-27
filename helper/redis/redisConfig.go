package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func GetRedisConnection() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("MYAPP_REDIS_HOST"),
		Password: os.Getenv("MYAPP_REDIS_PASS"),
		DB:       0,
	})
}
