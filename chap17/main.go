package main

import "fmt"

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    275,
}
var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func main() {
	list_17_2()
	list_17_5()
	list_17_6()
	list_17_7()
	list_17_8()
	list_17_9()
	list_17_11()
	list_17_12()
	list_17_13()
	list_17_14()
	list_17_15()
	list_17_16()
	list_17_17()
	list_17_18()
	list_17_19()
	list_17_20()
	list_17_21()
	list_17_22()

}

func list_17_2() {
	fmt.Println("======== list_17_2 ========")
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
}

func list_17_5() {
	fmt.Println("======== list_17_5 ========")
	// 각 출력항목마다 Space가 추가되고, 끝나면 개행문자가 온다.
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	// 각 항목모다 공백없이 연결되어 출력한다.(책에서는 설명이 좀 다름)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
}

func list_17_6() {
	fmt.Println("======== list_17_6 ========")
	// %v : 값을 일반적인 형태로 출력, 어떤 타입이든 기본적인 표현으로 보여준다.
	// %4.2f : 고정소수점 형식으로 소수점이하2자리, 최대 4자리로 출력한다.
	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
}

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		// 화면에 출력하는 대신 해당 결과를 문자열로 반환한다.
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		// 오류가 나오면 해당 오류메세지를 문자열로 반환한다.
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

func list_17_7() {
	fmt.Println("======== list_17_7 ========")
	name, _ := getProductName(1)
	fmt.Println(name)
	_, err := getProductName(10)
	fmt.Println(err.Error())

}

// 가변인자를 받아서 출력한다.
// Printfln("Hello %v", "World")
// Printfln("Number: %d, Float: %4.2f", 10, 3.14159)
func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func list_17_8() {
	fmt.Println("======== list_17_8 ========")
	// 기본포맷으로 출력한다. 구조체Kayak를 필드명없이 값만 순서대로 출력한다.
	Printfln("Value: %v", Kayak)
	// 구조체이름과 필드명 값 모두를 출력한다.
	Printfln("Go syntax: %#v", Kayak)
	// %T : 변수의 타입을 출력한다.
	Printfln("Type: %T", Kayak)
}

func list_17_9() {
	// 책의 설명과는 다른 결과가 나오는데 이유가 뭘까?
	fmt.Println("======== list_17_9 ========")
	// 구조체라면 값만 출력한다.
	Printfln("Value: %v", Kayak)
	// 구조체라면 필드 이름과 값이 출력된다.
	Printfln("Value with fields: %+v", Kayak)
}

// Product 의 메서드(리시버)
// Product구조체로 만들어진 변수는 String()함수를 호출할수있다.
func (p Product) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

func list_17_11() {
	fmt.Println("======== list_17_11 ========")
	number := 250
	// 250을 2진수, 10진수,8진수,16진수로 출력한다.
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)
}

func list_17_12() {
	fmt.Println("======== list_17_12 ========")
	number := 279.00
	// 부동소수점 숫자를 2진수 지수 형식으로 출력
	Printfln("Decimalless with exponent: %b", number)
	// 부동소수점을 지수 표기법 (scientific notation) 으로 출력
	Printfln("Decimal with exponent: %e", number)
	// 일반적인 소수점 표기로 출력
	Printfln("Decimal without exponent: %f", number)
	// 숫자를 16진수로 출력합니다. (주로 정수용임)
	Printfln("Hexadecimal: %x, %X", number, number)
}

func list_17_13() {
	fmt.Println("======== list_17_13 ========")
	number := 279.00
	// %8.2f : 부동소수점을 소수점 표기로 출력. 소숫점이하는 2자리, 최대길이 8자리, 결과는 오른쪽정렬
	Printfln("Decimal without exponent: >>%8.2f<<", number)
}

func list_17_14() {
	fmt.Println("======== list_17_14 ========")
	number := 279.00
	// 최대길이가 없기때문에 해당 출력의 길이만큼만 출력한다.
	Printfln("Decimal without exponent: >>%.2f<<", number)
}

func list_17_15() {
	fmt.Println("======== list_17_15 ========")
	number := 279.00
	// %+.2f : 항상 양수,또는 음수를 출력
	Printfln("Sign: >>%+.2f<<", number)
	// 빈공간이 생기면 공백대신 0으로 채운다.
	Printfln("Zeros for Padding: >>%010.2f<<", number)
	// 출력시 좌측정렬을 한다.
	Printfln("Right Padding: >>%-8.2f<<", number)
}

func list_17_16() {
	fmt.Println("======== list_17_16 ========")
	name := "Kayak"
	// 일반 문자열 출력
	Printfln("String: %s", name)
	// 첫번째 문자출력
	Printfln("Character: %c", []rune(name)[0])
	// 출력이 U+로 시작하고, 해당 문자의 유니코드 출력
	Printfln("Unicode: %U", []rune(name)[0])
}

func list_17_17() {
	fmt.Println("======== list_17_17 ========")
	name := "Kayak"
	// %t : 참이면 true, 거짓이면 false를 출력한다.
	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)
}

func list_17_18() {
	fmt.Println("======== list_17_18 ========")
	name := "Kayak"
	// 해당 포인터의 저장위치를 16진수로 출력한다
	Printfln("Pointer: %p", &name)
}

func list_17_19() {
	fmt.Println("======== list_17_19 ========")
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	// 이름 , 카테고리 , 가격을 입력받는다.
	n, err := fmt.Scan(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func list_17_20() {
	fmt.Println("======== list_17_20 ========")
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	// 한줄안에 모든 정보를 입력해야만 한다.
	n, err := fmt.Scanln(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func list_17_21() {
	fmt.Println("======== list_17_21 ========")
	var name string
	var category string
	var price float64
	source := "Lifejacket Watersports 48.95"
	// 기존 스캔함수와 동일하지만, 키보드 입력받는대신 soruce문자열로 입력을 받는다.
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func list_17_22() {
	fmt.Println("======== list_17_22 ========")
	var name string
	var category string
	var price float64
	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	// soruce문자열을 template포맷으로 변환하여 Sscanf문에 입력을 받는다.
	// 입력은 3개를 입력받아야 하는데 Product를 포함한 문자열은 4개인데, 문제 없나?
	// 굳이 Product를 추가한 이유는?
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
