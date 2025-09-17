package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Println("main function started")
		CalcStoreTotal(Products)
		fmt.Println("main function complete")
	*/

	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}
