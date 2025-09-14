package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			panic(message.CategoryError)
			//fmt.Println(message.Category, "(no such category)")
		}
	}
}
