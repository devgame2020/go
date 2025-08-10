package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"

	"gameserver/config"
	"gameserver/models"
	"gameserver/utils"
)

// 요청 DTO
type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GuestUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var ctx = context.Background()

// return
// token : token
// msg : message
// ret : true(성공), false(실패)
func _CreateUser(Username string, Password string, Logintype int) (token string, requestType int, ret error) {

	rdb := config.GetRedisClient()
	key := "user:" + Username
	// 사용자 존재 여부 확인
	exists, _ := rdb.Exists(ctx, key).Result()
	if exists == 1 {
		return "", http.StatusBadRequest, errors.New("user already exists")
	}

	// 비밀번호 해싱
	hashedPassword, err := utils.HashPassword([]byte(Password)) // bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", http.StatusBadRequest, errors.New("password hashing failed")
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Password hashing failed"})
	}

	// 고유 ID 생성 (예: UUID)
	userID := utils.GenerateUUID() // uuid.New().String()

	nickname := Username

	if Logintype == int(models.LoginTypeGuest) {
		nickname = "guest" + Username[:10]
	}

	// 유저 정보 Redis에 저장
	user := models.User{
		ID:        userID,
		Username:  Username,
		Password:  string(hashedPassword),
		Nickname:  nickname,
		LoginType: models.LoginTypeUser,
		Coins:     1000,
	}
	userData, _ := json.Marshal(user)

	err = rdb.HSet(ctx, key, "info", userData, "coin", 1000).Err()
	if err != nil {
		return "", http.StatusBadRequest, errors.New("failed to save user")
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save user"})
	}

	// JWT 토큰 생성
	// token, err := utils.GenerateToken(user.ID)
	token, err = utils.GenerateToken(user.Username)
	if err != nil {
		return token, http.StatusBadRequest, errors.New("failed to generate token")
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return token, http.StatusOK, nil
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"message": "User created",
	// 	"token":   token,
	// })
}

func _LoginUser(Username string, Password string) (token string, requestType int, ret error) {
	rdb := config.GetRedisClient()
	key := "user:" + Username
	// Redis에서 사용자 정보 조회
	userData, err := rdb.HGet(ctx, key, "info").Result()
	fmt.Println("==========userData=======")
	fmt.Println(userData)

	if err == redis.Nil {
		return token, http.StatusNotFound, errors.New("user not found")
		// return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
	} else if err != nil {
		return token, http.StatusInternalServerError, errors.New("redis error")
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Redis error"})
	}

	var user models.User
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return token, http.StatusBadRequest, errors.New("failed to parse user data")
		//return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to parse user data"})
	}

	fmt.Println("==========login=======")
	fmt.Println(user)

	coin, err := rdb.HGet(ctx, key, "coin").Result()
	if err == redis.Nil {
		return token, http.StatusBadRequest, errors.New("coin not found")
	} else if err != nil {
		return token, http.StatusBadRequest, errors.New("redis error")
	}

	user.Coins = utils.Atoi(coin)

	// 비밀번호 검증
	if err := utils.CompareHashAndPassword([]byte(user.Password), []byte(Password)); err != nil {
		return token, http.StatusUnauthorized, errors.New("invalid password")
		// return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid password"})
	}

	// JWT 토큰 생성
	token, err = utils.GenerateToken(user.Username)
	if err != nil {
		return token, http.StatusBadRequest, errors.New("failed to generate token")
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return token, http.StatusOK, nil
}

// SignUp godoc
// @Summary OK 회원가입
// @Description 회원가입 API입니다
// @Tags init
// @Accept  json
// @Produce  json
// @Param userRequest body UserRequest true "SignUp Body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /init/signup [post]
func SignUp(c echo.Context) error {
	fmt.Println("===SignUp====")

	// ================= 입력값 검사 ======================================
	var req UserRequest

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
	// ======================================================================

	token, erroType, err := _CreateUser(req.Username, req.Password, int(models.LoginTypeUser))

	if err != nil {
		return c.JSON(erroType, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created",
		"token":   token,
	})

	/*
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
	*/
}

// guestLogin godoc
// @Summary OK Guest로그인
// @Description Guest로그인 API입니다
// @Description 아이디를 입력하면 token을 반환합니다.
// @Tags init
// @Accept  json
// @Produce  json
// @Param userRequest body GuestUserRequest true "Login Body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /init/guest-login [post]
func GuestLogin(c echo.Context) error {
	var req GuestUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username is required"})
	}

	token := ""
	password := req.Username[:10]

	token, _, err := _LoginUser(req.Username, password)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Login successful",
			"token":   token,
		})
	}

	//
	token, erroType, err := _CreateUser(req.Username, password, int(models.LoginTypeGuest))
	if err != nil {
		return c.JSON(erroType, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created",
		"token":   token,
	})
}

// SignUp godoc
// @Summary OK 로그인
// @Description 로그인 API입니다
// @Description 아이디,패스워드를 입력하면 token을 반환합니다.
// @Tags init
// @Accept  json
// @Produce  json
// @Param userRequest body UserRequest true "Login Body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /init/login [post]
func Login(c echo.Context) error {
	var req UserRequest

	// =======================================================================================
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request."})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password are required"})
	}
	// =======================================================================================

	token, erroType, err := _LoginUser(req.Username, req.Password)

	if err != nil {
		return c.JSON(erroType, echo.Map{"error": err.Error()})
		// return c.JSON(http.StatusInternalServerError, echo.Map{"error": errorMsg})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
	})
}
