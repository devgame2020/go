package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// jwtKey는 JWT 서명용 비밀키
var jwtKey = []byte("your-secret-key")

// JWT Claims 구조
type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// func CreateJWT(userID string) (string, error) {
// 	// 유효기간 24시간 토큰을 생성한다.
// 	claims := &Claims{
// 		UserID: userID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtKey)
// }

// GenerateToken 함수는 userID를 받아 JWT 토큰을 생성합니다.
// 회원가입, 로그인시 호출된다.
func GenerateToken(userName string) (string, error) {
	fmt.Println("==GenerateToken===")
	fmt.Println(userName)
	claims := &Claims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24시간 유효
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken 함수는 JWT 토큰 문자열을 받아서 유효성 검증 후 클레임을 반환합니다.
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 서명 방식 체크 (HS256인지 확인)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing m ethod")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GetUserIDFromToken(c echo.Context) (string, error) {
	val := c.Get("username")
	user_id, ok := val.(string)

	fmt.Println("=======GetUserIDFromToken========")
	fmt.Println(val)
	fmt.Println(user_id, ok)

	if !ok {
		return "", errors.New("no user in context")
	}
	// user_id := c.Get("userID").(string)
	// if user_id == nil {
	// 	return "", errors.New("no user in context")
	// }
	// token := user.(*jwt.Token)
	// claims := token.Claims.(*Claims)
	return user_id, nil
}
