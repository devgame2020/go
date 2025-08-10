package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"gameserver/models"
	"gameserver/utils"
)

type UserProfileResponse struct {
	UserID    int    `json:"user_id" example:"123"`
	Username  string `json:"username" example:"john_doe"`
	Coins     int    `json:"coins" example:"1500"`
	LastLogin string `json:"last_login" example:"2023-10-01"`
}

// GetProfile godoc
// @Summary 유저 프로필 조회
// @Description Bearer 토큰을 통해 유저 정보를 조회합니다.
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} UserProfileResponse
// @Failure 401 {object} ErrorResponse
// @Router /user/profile [get]
func GetProfile(c echo.Context) error {
	fmt.Println("===get profile====")
	fmt.Println(c)

	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	fmt.Println("===get profile2====")

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	fmt.Println("===get profile3====")

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":    user.ID,
		"username":   user.Username,
		"coins":      user.Coins,
		"last_login": user.LastLogin,
	})
}

// ======== Nick Name 변경 =====================
// ==
func SetNickname(c echo.Context) error {
	var req struct {
		Nickname string `json:"nickname"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	fmt.Println("===get profile2====")

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	fmt.Println("===get profile3====")

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":    user.ID,
		"username":   user.Username,
		"coins":      user.Coins,
		"last_login": user.LastLogin,
	})
}
