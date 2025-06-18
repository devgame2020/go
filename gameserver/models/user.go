package models

import (
	"context"
	"errors"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type User struct {
	ID        string // Redis key에서 주로 string ID를 사용
	Username  string
	Password  string // 해시된 비밀번호 저장
	Coins     int
	LastLogin string // "YYYY-MM-DD"
}

// Redis 초기화 함수 (main.go에서 호출)
func InitRedis(addr, password string, db int) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(ctx).Result()
	return err
}

// 유저 정보 가져오기
func GetUser(userID string) (*User, error) {
	key := "user:" + userID
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, errors.New("user not found")
	}

	// Redis Hash에서 모든 필드 조회
	fields, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	coins, _ := strconv.Atoi(fields["coins"])

	user := &User{
		ID:        userID,
		Username:  fields["username"],
		Password:  fields["password"],
		Coins:     coins,
		LastLogin: fields["last_login"],
	}

	return user, nil
}

// 유저 정보 저장 (업데이트)
func SaveUser(user *User) error {
	key := "user:" + user.ID

	_, err := rdb.HSet(ctx, key, map[string]interface{}{
		"coins":      user.Coins,
		"last_login": user.LastLogin,
	}).Result()

	return err
}
