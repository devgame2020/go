package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	fmt.Println("======== list_11_30 ========")
	fmt.Println("== 슬라이스 expenses의 element 에는 Service, Product 둘다 저장되어있다. ==")
	fmt.Println("== expense.(Service)로 타입단언을 수행하여 성공하면 Service의 멤버값들을 사용하고, 실패하면 interface의 메소드를 호출한다. ==")

	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}
}
