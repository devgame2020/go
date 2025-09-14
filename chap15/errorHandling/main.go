package main

import "fmt"

func main() {
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
				panic(err)
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()

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
