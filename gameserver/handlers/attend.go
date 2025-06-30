package handlers

// 출석부관리
// 1주일 출석부
// 1달 출석부
// 신규유저출석부
// 복귀유저출석부

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"gameserver/models"
	"gameserver/utils"
)

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
