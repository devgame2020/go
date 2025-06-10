package main

import (
	"fmt"       // 기본 입출력 패키지
	"math/rand" // 난수생성 패키지
	"sort"
)

func main() {
	// f4_01()
	// f4_10()
	// f4_99()
	// f4_102()
	// f4_105()
	// f4_106()
	// f4_107()
	// f4_112()
	// f4_114()
	f4_115()
}

func f4_01() {
	// 표준 패키지 사용예시
	fmt.Println("Value:", rand.Int())

	// 리터럴 값
	fmt.Println("Hello", "Go")
	fmt.Println(20 + 30)

	// 상수 사용
	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println(price + tax)

	// 타입없는 상수
	const quantity = 2
	fmt.Println("Total:", quantity*(price+tax))
}

func f4_10() {
	const price, tax float32 = 275, 27.50
	const quantity, inStock = 2, true

	fmt.Println("=== 여러 상수 정의 ======")
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)

	// 상수 quantity를 자동으로 타입변환한다.
	fmt.Println("Total:", 2*quantity*(price+tax))

}

func f4_99() {
	var price float32 = 275.00
	var tax float32 = 27.50

	fmt.Println("=== 변수 사용 ======")
	fmt.Println(price + tax)
	price = 300
	fmt.Println(price + tax)

	fmt.Println("=== 변수 : 자료형 생략 ======")
	var price2 = 275.00
	var price3 = price2
	fmt.Println(price2)
	fmt.Println(price3)

	// 타입을 혼합하여 계산안됨
	// var price5 = 275.00
	// var tax2 float32 = 27.50
	// fmt.Println(price5 + tax2)

}

func f4_102() {
	// 변수값 초기화 생략
	var price float32
	fmt.Println(price)
	price = 275.00
	fmt.Println(price)

	// 한번에 여러변수정의
	var price2, tax2 = 275.00, 27.50
	fmt.Println(price2 + tax2)

	// 짧은 변수 선언
	price3 := 275.00
	fmt.Println(price3)

	price4, tax4, inStock4 := 275.00, 27.50, true
	fmt.Println("Total:", price4+tax4)
	fmt.Println("In Stock:", inStock4)
}

func f4_105() {
	price, tax, inStock := 275.00, 27.50, true
	fmt.Println("Total:", price+tax)
	fmt.Println("In Stock:", inStock)

	// 변수재정의
	// 적어도 하나가 이미 존재하지 않고, 변수타입을 변경하지 않는다면 재정의허용
	price2, tax := 200.0, 25.00
	fmt.Println("Total 2:", price2+tax)
}

func f4_106() {
	// 빈식별자 사용
	// Go는 변수를 정의하고 사용하지 않으면 오류가 나온다.
	price, tax, inStock, _ := 275.00, 27.50, true, true
	var _ = "Alice"
	fmt.Println("Total:", price+tax)
	fmt.Println("In Stock:", inStock)
}

func f4_107() {
	first := 100
	second := first
	first++
	fmt.Println("First:", first)
	fmt.Println("Second:", second)

	// 포인터 역참조
	first2 := 100
	var second2 *int = &first // 포인터변수 second2에 first의 주소값을 저장한다.
	second3 := &first         //

	first2++
	fmt.Println("First2:", first2)
	fmt.Println("Second2:", second2)
	fmt.Println("Second2:", *second2)
	fmt.Println("Second3:", *second3)

	var myPointer *int
	myPointer = second3
	*myPointer++

	fmt.Println("First:", first)
	fmt.Println("myPointer:", *myPointer)
}

func f4_112() {
	first := 100
	var second *int

	fmt.Println(second)
	second = &first
	// fmt.Println(second)
	fmt.Println(second == nil)
}

func f4_114() {
	first := 100
	second := &first
	third := &second

	fmt.Println(first)
	fmt.Println(*second)
	fmt.Println(**third)
}

func f4_115() {
	// secondName 은 names[1]의 value를 저장하고 있기때문에 sort를 하더라도 값이 변하지 않는다.
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	fmt.Println(secondName)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println(names)

	// secondPosition은 names[1]의 주소값을 저장하고 있기때문에 정렬후 names[1]의 주소에 해당하는 값이 변경될수있다.
	names2 := [3]string{"Alice", "Charlie", "Bob"}
	secondPosition := &names2[1]
	fmt.Println(*secondPosition)
	sort.Strings(names2[:])
	fmt.Println(*secondPosition)

}
