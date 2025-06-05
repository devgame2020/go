package main

import "fmt"

type User struct {
	Name, Email, Phone string
	WillAttend         bool
}

// 초기사이즈 : 0
// 용량 : 10
// [] : 슬라이스
var users = make([]*User, 0, 10)

func main() {
	fmt.Println("Example 01")

	// User 데이터 추가
	users = append(users, &User{
		Name:       "홍길동",
		Email:      "hong@example.com",
		Phone:      "010-1234-5678",
		WillAttend: true,
	})

	users = append(users, &User{
		Name:       "김철수",
		Email:      "kim@example.com",
		Phone:      "010-8765-4321",
		WillAttend: false,
	})

	// 출력
	for i, user := range users {
		fmt.Printf("User %d:\n", i+1)
		fmt.Printf("  Name: %s\n", user.Name)
		fmt.Printf("  Email: %s\n", user.Email)
		fmt.Printf("  Phone: %s\n", user.Phone)
		fmt.Printf("  Will Attend: %t\n\n", user.WillAttend)
	}
}
