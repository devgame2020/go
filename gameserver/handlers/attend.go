package handlers

// 출석부관리
// 1주일 출석부
// 1달 출석부
// 신규유저출석부
// 복귀유저출석부

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gameserver/config"
	"gameserver/models"
	"gameserver/utils"
)

// var ctx = context.Background()

// 출석부 현황조회
func GetAttend(c echo.Context) error {
	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(username)
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

// 출석부 보상받기
func RewardAttend() {

}

// 출석부 초기화하기
func InitAttend() {

}

func CheckAttendance(c echo.Context) error {
	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	// 오늘 날짜
	today := time.Now().Format("2006-01-02")

	// 마지막 출석일 확인
	lastDate, _ := config.Redis.Get(ctx, "attendance_date:"+username).Result()
	if lastDate == today {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Already checked in today"})
	}

	// 출석일수 가져오기
	dayStr, _ := config.Redis.Get(ctx, "attendance:"+username).Result()
	day := utils.Atoi(dayStr)

	// 다음 출석일 계산 (1~7 순환)
	day = (day % 7) + 1
	coins := day * 100

	// 유저 정보 업데이트
	newCoins := user.Coins + coins
	key := "user:" + username
	// Redis에 업데이트
	err = config.Redis.HSet(config.Ctx, key, "coins", newCoins).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "코인 업데이트 실패"})
	}

	// 출석 정보 업데이트
	config.Redis.Set(ctx, "attendance:"+username, day, 0)
	config.Redis.Set(ctx, "attendance_date:"+username, today, 0)

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("%d일차 출석 완료! %d 코인을 획득했습니다.", day, coins),
		"coins":   user.Coins,
	})
}

