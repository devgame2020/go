package main

import (
	"fmt"
	_ "packages/data"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("main start")

	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))
}
