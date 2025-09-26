package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	// Product 인스턴스
	laptop := Product{name: "Laptop", category: "Electronics", price: 1200.00}
	phone := Product{name: "Smartphone", category: "Electronics", price: 800.00}

	// Service 인스턴스
	cloud := Service{description: "Cloud Storage", durationMonths: 12, monthlyFee: 10.00}
	antivirus := Service{description: "Antivirus Protection", durationMonths: 6, monthlyFee: 5.00}

	// Expense 인터페이스로 묶기
	expenses2 := []Expense{laptop, phone, cloud, antivirus}

	// 모든 비용 출력
	for _, e := range expenses2 {
		fmt.Printf("Name: %-20s | Monthly/One-time Cost: $%.2f | Annual/Full Cost: $%.2f\n",
			e.getName(), e.getCost(false), e.getCost(true))
	}

	// kayak := Product{"Kayak", "Watersports", 275}
	// insurance := Service{"Boat Cover", 12, 89.50}
	// fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	// fmt.Println("Service:", insurance.description, "Price:",
	// 	insurance.monthlyFee*float64(insurance.durationMonths))
}
