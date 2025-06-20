package router

import (
	"echo/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("", handler.GetUsers)    // /users
	userGroup.GET("/:id", handler.GetUser) // /users/:id
	userGroup.POST("", handler.CreateUser)
	userGroup.PUT("/:id", handler.UpdateUser)
	userGroup.DELETE("/:id", handler.DeleteUser)
}
