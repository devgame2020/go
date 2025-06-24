package main

import "fmt"

func main() {
	list_8_4()
	list_8_5()
	list_8_6()
	list_8_7()
	list_8_8()
	list_8_9()
	list_8_10()
	list_8_11()
	list_8_12()
	list_8_13()
	list_8_14()
	list_8_15()
	list_8_16()
	list_8_17()
	list_8_18()
	list_8_19()
	list_8_20()
	list_8_21()
	list_8_22()
	list_8_23()
	list_8_24()

}

func printPrice4() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

func list_8_4() {
	fmt.Println("======== list_8_4 ========")
	fmt.Println("About to call function")
	printPrice4()
	fmt.Println("Function complete")
}

func printPrice5(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func list_8_5() {
	fmt.Println("======== list_8_5 ========")
	printPrice5("Kayak", 275, 0.2)
	printPrice5("Lifejacket", 48.95, 0.2)
	printPrice5("Soccer Ball", 19.50, 0.15)
}

func printPrice6(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func list_8_6() {
	fmt.Println("======== list_8_6 ========")
	printPrice6("Kayak", 275, 0.2)
	printPrice6("Lifejacket", 48.95, 0.2)
	printPrice6("Soccer Ball", 19.50, 0.15)
}

func printPrice7(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func list_8_7() {
	fmt.Println("======== list_8_7 ========")
	printPrice7("Kayak", 275, 0.2)
	printPrice7("Lifejacket", 48.95, 0.2)
	printPrice7("Soccer Ball", 19.50, 0.15)
}

func printPrice8(string, float64, float64) {
	// taxAmount := price * 0.25
	fmt.Println("No parameters")
}

func list_8_8() {
	fmt.Println("======== list_8_8 ========")
	printPrice8("Kayak", 275, 0.2)
	printPrice8("Lifejacket", 48.95, 0.2)
	printPrice8("Soccer Ball", 19.50, 0.15)
}

func printSuppliers9(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func list_8_9() {
	fmt.Println("======== list_8_9 ========")
	printSuppliers9("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers9("Lifejacket", []string{"Sail Safe Co"})
}

func printSuppliers10(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func list_8_10() {
	fmt.Println("======== list_8_10 ========")
	printSuppliers10("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers10("Lifejacket", "Sail Safe Co")
}

func printSuppliers11(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func list_8_11() {
	fmt.Println("======== list_8_11 ========")
	printSuppliers11("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers11("Lifejacket", "Sail Safe Co")
	printSuppliers11("Soccer Ball")
}

func printSuppliers12(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func list_8_12() {
	fmt.Println("======== list_8_12 ========")
	printSuppliers12("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers12("Lifejacket", "Sail Safe Co")
	printSuppliers12("Soccer Ball")
}

func printSuppliers13(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func list_8_13() {
	fmt.Println("======== list_8_13 ========")
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers13("Kayak", names...)
	printSuppliers13("Lifejacket", "Sail Safe Co")
	printSuppliers13("Soccer Ball")
}

func swapValues14(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

func list_8_14() {
	fmt.Println("======== list_8_14 ========")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues14(val1, val2)
	fmt.Println("After calling function", val1, val2)
}

func swapValues15(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func list_8_15() {
	fmt.Println("======== list_8_15 ========")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues15(&val1, &val2)
	fmt.Println("After calling function", val1, val2)
}

func calcTax16(price float64) float64 {
	return price + (price * 0.2)
}

func list_8_16() {
	fmt.Println("======== list_8_16 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		priceWithTax := calcTax16(price)
		fmt.Println("Product: ", product, "Price:", priceWithTax)
	}
}

func calcTax17(price float64) float64 {
	return price + (price * 0.2)
}

func list_8_17() {
	fmt.Println("======== list_8_17 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		fmt.Println("Product: ", product, "Price:", calcTax17(price))
	}
}

func swapValues18(first, second int) (int, int) {
	return second, first
}

func list_8_18() {
	fmt.Println("======== list_8_18 ========")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	val1, val2 = swapValues18(val1, val2)
	fmt.Println("After calling function", val1, val2)
}

func calcTax19(price float64) float64 {
	if price > 100 {
		return price * 0.2
	}
	return -1
}

func list_8_19() {
	fmt.Println("======== list_8_19 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		tax := calcTax19(price)
		if tax != -1 {
			fmt.Println("Product: ", product, "Tax:", tax)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}
}

func calcTax20(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func list_8_20() {
	fmt.Println("======== list_8_20 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		taxAmount, taxDue := calcTax20(price)
		if taxDue {
			fmt.Println("Product: ", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}
}

func calcTax21(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func list_8_21() {
	fmt.Println("======== list_8_21 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		if taxAmount, taxDue := calcTax21(price); taxDue {
			fmt.Println("Product: ", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}
}

func calcTax22(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice22(products map[string]float64,
	minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax22(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func list_8_22() {
	fmt.Println("======== list_8_22 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	total1, tax1 := calcTotalPrice22(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice22(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
}

func calcTotalPrice23(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}
	return
}

func list_8_23() {
	fmt.Println("======== list_8_23 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	_, total := calcTotalPrice23(products)
	fmt.Println("Total:", total)
}

func calcTotalPrice24(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func list_8_24() {
	fmt.Println("======== list_8_24 ========")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	_, total := calcTotalPrice24(products)
	fmt.Println("Total:", total)
}
