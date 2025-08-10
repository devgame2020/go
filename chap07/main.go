package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	var products []string

	fmt.Println(products)
	fmt.Println("len and cap ", len(products), cap(products))
	products = append(products, "a")
	products = append(products, "b")
	products = append(products, "c")
	fmt.Println(products)
	fmt.Println("len and cap ", len(products), cap(products))
	// fmt.Println("len and cap ", len(someNames), cap(someNames))
	// someNames = append(someNames, "Gloves")
	// fmt.Println("len and cap ", len(someNames), cap(someNames))
	// copy(someNames, allNames)
	// fmt.Println("someNames:", someNames)
	// fmt.Println("allNames", allNames)

	list_7_4()
	list_7_5()
	list_7_6()
	list_7_7()
	list_7_8()
	list_7_9()
	list_7_10()
	list_7_11()
	list_7_12()
	list_7_13()
	list_7_14()
	list_7_15()
	list_7_16()
	list_7_17()
	list_7_18()
	list_7_19()
	list_7_20()
	list_7_21()
	list_7_22()
	list_7_23()
	list_7_24()
	list_7_25()
	list_7_26()
	list_7_27()
	list_7_28()
	list_7_29()
	list_7_30()
	list_7_31()
	list_7_32()
	list_7_33()
	list_7_34()
	list_7_35()
	list_7_36()
	list_7_37()
	list_7_38()
	list_7_39()
	list_7_40()
	list_7_41()
	list_7_42()
	list_7_43()
	list_7_44()
	list_7_45()
	list_7_46()
	list_7_47()
	list_7_48()
	list_7_49()
	list_7_50()
	list_7_rune()

}

func list_7_4() {
	fmt.Println("======== list_7_4 ========")
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
}

func list_7_5() {
	fmt.Println("======== list_7_5 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
}

func list_7_6() {
	fmt.Println("======== list_7_6 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// 아래 구문은 오류. names와 길이가 정확하게 일치해야한다.
	// var otherArray [4]string = names
	fmt.Println(names)
}

func list_7_7() {
	fmt.Println("======== list_7_7 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := names
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", otherArray)
}

func list_7_8() {
	fmt.Println("======== list_7_8 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := &names
	names[0] = "Canoe"
	otherArray[1] = "abc"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", *otherArray)
}

func list_7_9() {
	fmt.Println("======== list_7_9 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := names == moreNames
	fmt.Println("comparison:", same)
}

func list_7_10() {
	fmt.Println("======== list_7_10 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func list_7_11() {
	fmt.Println("======== list_7_11 ========")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, value := range names {
		fmt.Println("Value:", value)
	}
}

func list_7_12() {
	fmt.Println("======== list_7_12 ========")
	names := make([]string, 3)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
}

func list_7_13() {
	fmt.Println("======== list_7_13 ========")
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
}

func list_7_14() {
	fmt.Println("======== list_7_14 ========")
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)
	fmt.Println("len and cap ", len(names), cap(names))
}

func list_7_15() {
	fmt.Println("======== list_7_15 ========")
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

func list_7_16() {
	fmt.Println("======== list_7_16 ========")
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
}

func list_7_17() {
	fmt.Println("======== list_7_17 ========")
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

func list_7_18() {
	fmt.Println("======== list_7_18 ========")
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	moreNames := []string{"Hat Gloves"}
	appendedNames := append(names, moreNames...)
	fmt.Println("appendedNames:", appendedNames)
}

func list_7_19() {
	fmt.Println("======== list_7_19 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func list_7_20() {
	fmt.Println("======== list_7_20 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func list_7_21() {
	fmt.Println("======== list_7_21 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")

	// cap을 초과하면 새로운 공간에 할당된다.
	// someNames = append(someNames, "Gloves", "aaa", "bbb")

	// someNames = append(someNames, "Gloves")
	// someNames = append(someNames, "aaa")
	// someNames = append(someNames, "bbb")

	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func list_7_22() {
	fmt.Println("======== list_7_22 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func list_7_23() {
	fmt.Println("======== list_7_23 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3:3]
	allNames := products[:]
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	someNames = append(someNames, "Gloves")
	//someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func list_7_24() {
	fmt.Println("======== list_7_24 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]   // products의 1번~끝까지의 Data로 Slice를 생성한다.
	someNames := allNames[1:3] // allNames의 1~2까지의 Data로 Slice를 생성한다.

	fmt.Println("=========================")
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)

	fmt.Println("===========append(allNames, 'Gloves')=============")
	allNames = append(allNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)

	fmt.Println("==========allNames[1] = 'Canoe'===============")
	allNames[1] = "Canoe"
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func list_7_25() {
	fmt.Println("======== list_7_25 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	allNames := products[1:]
	someNames := make([]string, 2)
	// someNames, allNames 는 각각의 배열을 갖는다.
	// allNames가 더 크더라도, someNames의 크기만큼만 복사한다.
	copy(someNames, allNames)

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func list_7_26() {
	fmt.Println("======== list_7_26 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	var someNames []string
	// fmt.Println("len and cap ", len(someNames), cap(someNames))
	// someNames = append(someNames, "Gloves")
	// fmt.Println("len and cap ", len(someNames), cap(someNames))
	copy(someNames, allNames)
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func list_7_27() {
	fmt.Println("======== list_7_27 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := []string{"Boots", "Canoe"}
	copy(someNames[1:], allNames[2:3])
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)

	products2 := [4]string{"0", "1", "2", "3"}
	fmt.Println("products2", products2[3:3])

}

func list_7_28() {
	fmt.Println("======== list_7_28 ========")
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products, replacementProducts)
	fmt.Println("products:", products)
}

func list_7_29() {
	fmt.Println("======== list_7_29 ========")
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products[0:1], replacementProducts)
	fmt.Println("products:", products)
}

func list_7_30() {
	fmt.Println("======== list_7_30 ========")
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products[:2], products[3:]...)
	fmt.Println("Deleted:", deleted)
	fmt.Println("products:", products)
	fmt.Println("len and cap ", len(deleted), cap(deleted))
}

func list_7_31() {
	fmt.Println("======== list_7_31 ========")
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func list_7_32() {
	fmt.Println("======== list_7_32 ========")
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	sort.Strings(products)
	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func list_7_33() {
	fmt.Println("======== list_7_33 ========")
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	// fmt.Println("Equal:", p1 == p2)

	fmt.Println(p1)
	fmt.Println(p2)
}

func list_7_34() {
	fmt.Println("======== list_7_34 ========")
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
}

func list_7_35() {
	fmt.Println("======== list_7_35 ========")
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println(array)
}

func list_7_36() {
	fmt.Println("======== list_7_36 ========")
	products := make(map[string]float64, 10)
	products["Kayak"] = 279
	products["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])
}

func list_7_37() {
	fmt.Println("======== list_7_37 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])
}

func list_7_38() {
	fmt.Println("======== list_7_38 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	fmt.Println("Hat:", products["Hat"])
}

func list_7_39() {
	fmt.Println("======== list_7_39 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	value, ok := products["Hat"]
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

func list_7_40() {
	fmt.Println("======== list_7_40 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

func list_7_41() {
	fmt.Println("======== list_7_41 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	delete(products, "Hat")
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

func list_7_42() {
	fmt.Println("======== list_7_42 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func list_7_43() {
	fmt.Println("======== list_7_43 ========")
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}
}

func list_7_44() {
	fmt.Println("======== list_7_44 ========")
	var price string = "$48.95"
	var currency byte = price[0]
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func list_7_45() {
	fmt.Println("======== list_7_45 ========")
	var price string = "$48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func list_7_46() {
	fmt.Println("======== list_7_46 ========")
	// var price string = "€48.95"
	// var currency string = string(price[0])
	// var amountString string = price[1:]
	// amount, parseErr := strconv.ParseFloat(amountString, 64)
	// fmt.Println("Currency:", currency)
	// if parseErr == nil {
	// 	fmt.Println("Amount:", amount)
	// } else {
	// 	fmt.Println("Parse Error:", parseErr)
	// }
}

func list_7_47() {
	fmt.Println("======== list_7_47 ========")
	var price string = "€48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func list_7_48() {
	fmt.Println("======== list_7_48 ========")
	var price []rune = []rune("€48.95")
	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func list_7_49() {
	fmt.Println("======== list_7_49 ========")
	var price = "€48.95"
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
}

func list_7_50() {
	fmt.Println("======== list_7_50 ========")
	var price = "€48.95"
	for index, char := range []byte(price) {
		fmt.Println(index, char)
	}
}

func list_7_rune() {
	fmt.Println("======== list_7_rune ========")
	var ch rune = '한'
	fmt.Println(ch)        // 출력: 54620 (유니코드 코드포인트)
	fmt.Printf("%c\n", ch) // 출력: 한
}
