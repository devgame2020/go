package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err := Redis.Ping(Ctx).Result()
	return err
}

func GetRedisClient() *redis.Client {
	if Redis == nil {
		InitRedis()
	}
	return Redis
}

func CloseRedis() {
	if Redis != nil {
		_ = Redis.Close()
	}
}
