package routes

import (
	"gameserver/handlers"
	"gameserver/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// 회원가입 & 로그인 - 인증 필요 없음
	user := e.Group("/user")
	user.POST("/signup", handlers.SignUp)
	user.POST("/login", handlers.Login)

	// 인증 필요한 그룹
	game := e.Group("/game")
	game.Use(middleware.JWTAuthMiddleware())
	game.GET("/profile", handlers.GetProfile)

	game.POST("/roll", handlers.RollDice)           // 주사위 던지기
	game.GET("/daily-reward", handlers.DailyReward) // 일일보상
}
