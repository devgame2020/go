package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"gameserver/config"
	"gameserver/models"
	"gameserver/utils"
)

// 요청 DTO
type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var ctx = context.Background()

// 회원가입
func SignUp(c echo.Context) error {
	fmt.Println("===SignUp====")

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		log.Println("Bind error:", err)

		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password are required"})
	}

	if strings.HasPrefix(strings.ToLower(req.Username), "guest") {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username cannot start with 'guest'"})
	}

	// Redis 클라이언트
	rdb := config.GetRedisClient()

	// 사용자 존재 여부 확인
	exists, err := rdb.Exists(ctx, "user:"+req.Username).Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to check user existence"})
	}
	if exists == 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "User already exists"})
	}

	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Password hashing failed"})
	}

	// 고유 ID 생성 (예: UUID)
	userID := uuid.New().String()

	// 유저 정보 Redis에 저장
	user := models.User{
		ID:        userID,
		Username:  req.Username,
		Password:  string(hashedPassword),
		Nickname:  req.Username,
		LoginType: models.LoginTypeUser,
		Coins:     1000,
	}

	userData, _ := json.Marshal(user)
	key := "user:" + req.Username
	err = rdb.HSet(ctx, key, "info", userData, "coin", 1000).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save user"})
	}

	// JWT 토큰 생성
	// token, err := utils.GenerateToken(user.ID)
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created",
		"token":   token,
	})
}

// 로그인
func GuestLogin(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// ======== guest process ===========
	// 해당 부분 작업해야함
	token := "aaaaaaaaaaaaaa"

	// ================================================

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// 로그인
func Login(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password are required"})
	}

	rdb := config.GetRedisClient()

	key := "user:" + req.Username
	// Redis에서 사용자 정보 조회
	userData, err := rdb.HGet(ctx, key, "info").Result()
	fmt.Println("==========userData=======")
	fmt.Println(userData)

	if err == redis.Nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Redis error"})
	}

	var user models.User
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse user data"})
	}

	fmt.Println("==========login=======")
	fmt.Println(user)

	coin, err := rdb.HGet(ctx, key, "coin").Result()
	user.Coins = utils.Atoi(coin)

	// 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid password"})
	}

	// JWT 토큰 생성
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}
