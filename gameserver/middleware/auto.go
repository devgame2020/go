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

		// 예: userID를 claim에서 추출하고 context에 저장
		// userIDFloat, ok := claims["user_id"].(float64)

		/*
			token := strings.TrimPrefix(auth, "Bearer ")
			fmt.Println(auth)
			fmt.Println(token)
			claims, err := utils.ParseToken(token)
			userID := claims.UserID
			if err != nil {
				fmt.Println("error", "Invalid token")
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}

			// 컨텍스트에 userID 저장
			fmt.Println("AuthMiddleware:userId:", userID)
			c.Set("userID", userID)
		*/

		return next(c)
	}
}

/*
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
*/
