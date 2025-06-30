package utils

import (
	"strconv"
)

// 문자열을 int로 변환하고, 실패하면 0을 반환
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
