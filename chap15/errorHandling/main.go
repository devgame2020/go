package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Running"}

	for _, cat := range categories {
		// Category가 Watersports, Chess 인 항목에 대한 합을 구한다.
		total, err := Products.TotalPrice(cat)

		if err == nil {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}
}
