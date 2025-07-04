package routes

import (
	"gameserver/handlers"
	"gameserver/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.CustomLogger)

	// 인증 필요없음
	RegisterRouteInit(e)
	RegisterRouteEtc(e)

	RegisterRouteUser(e)
	RegisterRoutePlay(e)

}

func RegisterRouteEtc(e *echo.Echo) {
	e.GET("/chk", handlers.CheckHandler)
}

func RegisterRouteInit(e *echo.Echo) {
	// 회원가입 & 로그인 - 인증 필요 없음
	init := e.Group("/init")
	init.POST("/signup", handlers.SignUp)
	init.POST("/login", handlers.Login)
	init.POST("/guest-login", handlers.GuestLogin) // guest login : uuid값을 사용하여 회원가입 | 로그인처리할것
}

func RegisterRouteUser(e *echo.Echo) {
	user := e.Group("/user")
	user.Use(middleware.AuthMiddleware)

	user.GET("/profile", handlers.GetProfile)
	user.GET("/nickname", handlers.SetNickname) // 닉네임변경(회원가입시 기본닉네임설정됨)

}

func RegisterRoutePlay(e *echo.Echo) {
	game := e.Group("/game")
	game.Use(middleware.AuthMiddleware)
	game.POST("/roll", handlers.RollDice)             // 주사위 던지기
	game.POST("/doubleroll", handlers.DoubleRollDice) // 주사위 던지기
	game.GET("/daily-reward", handlers.DailyReward)   // 일일보상
}

// 출석부
func RegisterRouteAttend(e *echo.Echo) {

}

// 일일 퀘스트
func RegisterRouteDailyQuest(e *echo.Echo) {

}

// 메세지함
func RegisterRouteMessage(e *echo.Echo) {

}

// 인벤토리
func RegisterRouteInventory(e *echo.Echo) {

}

