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

	/*
		Redis = redis.NewClient(&redis.Options{
			// 나중에 설정파일로 분리할것
			//		Addr:     os.Getenv("REDIS_ADDR"),     // 예: "localhost:6379"
			//		Password: os.Getenv("REDIS_PASSWORD"), // 없으면 ""
			Addr:     "redis-14753.c340.ap-northeast-2-1.ec2.redns.redis-cloud.com:14753", // 예: "localhost:6379"
			Password: "p7kMcM9vqkqRCsGiiNZSgBaAUdnNXQwj",                                  // 없으면 ""
			DB:       0,
		})

		err := Redis.Ping(Ctx).Err()
		if err != nil {
			log.Fatalf("❌ Redis 연결 실패: %v", err)
			return err
		}

		log.Println("✅ Redis 연결 성공")
		return err
	*/
}

func CloseRedis() {
	if Redis != nil {
		_ = Redis.Close()
	}
}
