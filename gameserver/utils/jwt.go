package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// jwtKey는 JWT 서명용 비밀키
var jwtKey = []byte("your-secret-key")

// JWT Claims 구조
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateJWT(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// GenerateToken 함수는 userID를 받아 JWT 토큰을 생성합니다.
func GenerateToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
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
			return nil, errors.New("unexpected signing method")
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

// 문자열을 int로 변환하고, 실패하면 0을 반환
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func GetUserIDFromToken(c echo.Context) (string, error) {
	user := c.Get("user")
	if user == nil {
		return "", errors.New("no user in context")
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(*Claims)
	return claims.UserID, nil
}
