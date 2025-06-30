package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gameserver/config"
	"gameserver/utils"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type LOGINTYPE int

const (
	LoginTypeGuest  LOGINTYPE = iota // 0 => guest login
	LoginTypeUser                    // 1 => id,password login
	LoginTypeGoogle                  // 2 => google login
)

type User struct {
	ID        string // Redis key에서 주로 string ID를 사용
	Username  string
	Password  string    // 해시된 비밀번호 저장
	Nickname  string    // 유저닉네임
	LoginType LOGINTYPE // Login type (member,guest, google,facebook,apple등등)
	Coins     int
	LastLogin string // "YYYY-MM-DD"
}

// Redis 초기화 함수 (main.go에서 호출)
// func InitRedis(addr, password string, db int) error {

// 	rdb = redis.NewClient(&redis.Options{
// 		Addr:     addr,
// 		Password: password,
// 		DB:       db,
// 	})

// 	_, err := rdb.Ping(ctx).Result()
// 	return err
// }

// 유저 정보 가져오기
func GetUser(username string) (*User, error) {
	rdb := config.GetRedisClient()

	key := "user:" + username
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("rdb exists")

	if exists == 0 {
		fmt.Println("user not found", username)
		return nil, errors.New("user not found")
	}

	fields, err := rdb.HGet(ctx, "user:"+username, "info").Result()

	fmt.Println("=========fields=========")
	fmt.Println(fields)
	// Redis Hash에서 모든 필드 조회
	// fields, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal([]byte(fields), &user); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	coin, err := rdb.HGet(ctx, "user:"+username, "coin").Result()
	user.Coins = utils.Atoi(coin)

	fmt.Println("=========user=========")
	fmt.Println(user)

	// coins, _ := strconv.Atoi(fields["coins"])

	// user := &User{
	// 	ID:        fields["id"],
	// 	Username:  fields["username"],
	// 	Password:  fields["password"],
	// 	Coins:     coins,
	// 	LastLogin: fields["last_login"],
	// }

	return &user, nil
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
