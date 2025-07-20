package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(expenses))

	// kayak := Product{"Kayak", "Watersports", 275}
	// insurance := Service{"Boat Cover", 12, 89.50}
	// fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	// fmt.Println("Service:", insurance.description, "Price:",
	// 	insurance.monthlyFee*float64(insurance.durationMonths))
}
