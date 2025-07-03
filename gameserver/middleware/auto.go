package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("===AuthMiddleware===")
		auth := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			fmt.Println("error: Missing or invalid token")
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing or invalid token"})
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")
		fmt.Println("===========tokenString=============")
		fmt.Println(tokenString)
		fmt.Println(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("===unexpected signing method=====")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		fmt.Println("===========token=============")
		fmt.Println(token)

		if err != nil || !token.Valid {
			fmt.Println("=======error0=======")
			fmt.Println(err)
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("=======error1=======")
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token claims"})
		}
		fmt.Println("==========claims===========")
		fmt.Println(claims)

		// 컨텍스트에 userID 저장
		fmt.Println("AuthMiddleware:userId:", claims["user_id"])
		c.Set("username", claims["username"])


		return next(c)
	}
}

