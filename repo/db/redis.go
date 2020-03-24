package db

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var redisCli *redis.Client

// InitRedisClient initialization redis client
func InitRedisClient() {
	redisCli = newRedisClient()
}

// newRedisClient new redis client
func newRedisClient() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect failed:%v", err)
		return nil
	}
	return redisCli
}

// GetRedisClient get global redisCli
func GetRedisClient() *redis.Client {
	return redisCli
}