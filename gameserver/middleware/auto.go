package middleware

import (
	"gameserver/utils"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid token"})
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		userID, err := utils.ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		// 컨텍스트에 userID 저장
		c.Set("userID", userID)

		return next(c)
	}
}

func JWTAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid token"})
			}

			tokenString := strings.TrimPrefix(auth, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token claims"})
			}

			// 예: userID를 claim에서 추출하고 context에 저장
			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid user_id in token"})
			}
			c.Set("userID", int(userIDFloat))

			return next(c)
		}
	}
}
