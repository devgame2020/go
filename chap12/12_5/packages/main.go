package main

import (
	"chap12/12_5/packages/store"
	"fmt"
)

func main() {
	product := store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
}
