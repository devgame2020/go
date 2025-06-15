package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("==== Chap 06 =======")

	list_6_2()
	list_6_4()
	list_6_5()
	list_6_6()
	list_6_7()
	list_6_8()
	list_6_9()
	list_6_10()
	list_6_11()
	list_6_12()
	list_6_13()
	list_6_14()
	list_6_15()
	list_6_16()
	list_6_17()
	list_6_18()
	list_6_19()
	list_6_20()
	list_6_21()
	list_6_22()
	list_6_23()
	list_6_24()
	list_6_25()
	list_6_26()
	list_6_27()

}

func list_6_2() {
	fmt.Println("======== list_6_2 ========")
	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)
}

func list_6_4() {
	fmt.Println("======== list_6_4 ========")
	kayakPrice := 275.00
	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}
}

func list_6_5() {
	// fmt.Println("======== list_6_5 ========")
	// kayakPrice := 275.00
	// if (kayakPrice > 100) {
	// 	fmt.Println("Price is greater than 100")
	// }
}

func list_6_6() {
	fmt.Println("======== list_6_6 ========")
	kayakPrice := 275.00
	if kayakPrice > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice < 300 {
		fmt.Println("Price is less than 300")
	}
}

func list_6_7() {
	fmt.Println("======== list_6_7 ========")
	kayakPrice := 275.00
	if kayakPrice > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice < 100 {
		fmt.Println("Price is less than 100")
	} else if kayakPrice > 200 && kayakPrice < 300 {
		fmt.Println("Price is between 200 and 300")
	}
}

func list_6_8() {
	fmt.Println("======== list_6_8 ========")
	kayakPrice := 275.00
	if kayakPrice > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice < 100 {
		fmt.Println("Price is less than 100")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}
}

func list_6_9() {
	fmt.Println("======== list_6_9 ========")
	kayakPrice := 275.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}
}

func list_6_10() {
	fmt.Println("======== list_6_10 ========")
	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}
}

func list_6_11() {
	fmt.Println("======== list_6_11 ========")
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}
}

func list_6_12() {
	fmt.Println("======== list_6_12 ========")
	counter := 0
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
		// if (counter > 3) {
		//
		// break
		// }
	}
}

func list_6_13() {
	fmt.Println("======== list_6_13 ========")
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
	}
}

func list_6_14() {
	fmt.Println("======== list_6_14 ========")
	for counter := 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
}

func list_6_15() {
	fmt.Println("======== list_6_15 ========")
	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}
}

func list_6_16() {
	fmt.Println("======== list_6_16 ========")
	product := "Kayak"
	for index := range product {
		fmt.Println("Index:", index)
	}
}

func list_6_17() {
	fmt.Println("======== list_6_17 ========")
	product := "Kayak"
	for _, character := range product {
		fmt.Println("Character:", string(character))
	}
}

func list_6_18() {
	fmt.Println("======== list_6_18 ========")
	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}
}

func list_6_19() {
	fmt.Println("======== list_6_19 ========")
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func list_6_20() {
	fmt.Println("======== list_6_20 ========")
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			fmt.Println("K or k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func list_6_21() {
	fmt.Println("======== list_6_21 ========")
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func list_6_22() {
	fmt.Println("======== list_6_22 ========")
	product := "Kayak"
	for index, character := range product {
		fmt.Println("=======", string(character), "======")
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func list_6_23() {
	fmt.Println("======== list_6_23 ========")
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}
}

func list_6_24() {
	fmt.Println("======== list_6_24 ========")
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", counter/2)
		default:
			fmt.Println("Non-prime value:", counter/2)
		}
	}
}

func list_6_25() {
	fmt.Println("======== list_6_25 ========")
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}
}

func list_6_26() {
	fmt.Println("======== list_6_26 ========")
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
}

func list_6_27() {
	fmt.Println("======== list_6_27 ========")
	counter := 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
}
