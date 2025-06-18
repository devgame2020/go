package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gameserver/config"
	"gameserver/models"
	"gameserver/utils"
)

// 주사위 굴리기 핸들러
func RollDice(c echo.Context) error {
	// 사용자 ID 가져오기 (AuthMiddleware에서 context에 저장됨)
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "토큰 오류"})
	}

	// Redis에서 사용자 정보 가져오기
	key := fmt.Sprintf("user:%d", userID)
	coinsStr, err := config.Redis.HGet(config.Ctx, key, "coins").Result()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "사용자 정보 조회 실패"})
	}

	coins := utils.Atoi(coinsStr)
	if coins < 30 {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "코인이 부족합니다"})
	}

	// 주사위 굴리기 (1~6)
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1

	// 보상 계산
	reward := dice * 10
	newCoins := coins - 30 + reward

	// Redis에 업데이트
	err = config.Redis.HSet(config.Ctx, key, "coins", newCoins).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "코인 업데이트 실패"})
	}

	// 결과 응답
	return c.JSON(http.StatusOK, echo.Map{
		"message": "주사위 굴리기 성공",
		"dice":    dice,
		"reward":  reward,
		"coins":   newCoins,
	})
}

func DailyReward(c echo.Context) error {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	// 오늘 날짜 문자열
	today := time.Now().Format("2006-01-02")

	if user.LastLogin == today {
		// 이미 오늘 보상 받음
		return c.JSON(http.StatusOK, echo.Map{
			"message": "오늘 이미 보상을 받았습니다",
			"coins":   user.Coins,
		})
	}

	// 보상 지급
	const reward = 100
	user.Coins += reward
	user.LastLogin = today

	err = models.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "보상 저장 실패"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "일일 보상을 받았습니다",
		"reward":  reward,
		"coins":   user.Coins,
	})
}

func GetProfile(c echo.Context) error {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":    user.ID,
		"username":   user.Username,
		"coins":      user.Coins,
		"last_login": user.LastLogin,
	})
}
