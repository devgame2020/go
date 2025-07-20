package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	fmt.Println("======== list_11_25 ========")
	fmt.Println("== var expense Expense = product : 오류발생 ==")
	fmt.Println("== var expense Expense = &product : 정상실행 ==")

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
}
