package main

import (
	"fmt"       // 기본 입출력 패키지
	"math/rand" // 난수생성 패키지
)

func main() {
	f4_01()
	f4_10()
	f4_99()
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
