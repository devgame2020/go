package main

import (
	"echo/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
