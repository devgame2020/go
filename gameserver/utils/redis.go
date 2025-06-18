package utils

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var (
	redisClient *redis.Client
	once        sync.Once
)

// InitRedis 초기화 함수 (앱 시작 시 호출)
func InitRedis() {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis 주소
			Password: "",               // 비밀번호 없으면 ""
			DB:       0,                // 기본 DB
		})

		// 연결 테스트
		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("❌ Redis 연결 실패: %v", err)
		} else {
			log.Println("✅ Redis 연결 성공")
		}
	})
}

// GetRedisClient는 전역 Redis 클라이언트를 반환합니다.
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient
}
