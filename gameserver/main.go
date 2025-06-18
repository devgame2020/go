package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gameserver/config"
	"gameserver/routes"
)

func main() {
	// 1. 환경 변수 로드
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env 파일을 찾을 수 없습니다. 시스템 환경 변수를 사용합니다.")
	}

	// 2. Redis 초기화
	if err := config.InitRedis(); err != nil {
		log.Fatalf("❌ Redis 초기화 실패: %v", err)
	}
	defer config.CloseRedis()

	// 3. Echo 서버 생성
	e := echo.New()

	// 4. CORS 및 로깅 미들웨어 추가
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 5. 라우트 등록
	routes.RegisterRoutes(e)

	// 6. 서버 실행
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 서버 실행 중: http://localhost:%s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
