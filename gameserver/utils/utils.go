package utils

import (
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// 문자열을 int로 변환하고, 실패하면 0을 반환
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func HashPassword(Password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	return hashedPassword, err
}

func CompareHashAndPassword(UserPassword []byte, ReqPassword []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(UserPassword), []byte(ReqPassword))
}

func GenerateUUID() string {
	return uuid.New().String()
}

