package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	fmt.Println("======== list_11_28 ========")
	fmt.Println("== s := expense.(Service) 를 사용하여 타입단언을 하였다.  ==")
	fmt.Println("== 타입단언은 인터페이스만 적용가능하고, 인터페이스값이 특정 동적 타입을 갖고 있음을 컴파일러에게 알려주기 위해 사용한다. ==")
	fmt.Println("== 타입변환은 인터페이스가 아닌 특정타입에만 적용할수있고, 동일한 필드가 있는 구조체 타입간의 변환과 같이 해당 타입의 구조가 호환되는경우만 사용가능 ==")

	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:",
			s.monthlyFee*float64(s.durationMonths))
	}
}
