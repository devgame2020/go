package main

import "fmt"

func main() {
	list_9_4()
	list_9_5()
	list_9_6()
	list_9_7()
	list_9_8()
	list_9_9()
	list_9_10()
	list_9_11()
	list_9_12()
	list_9_13()
	list_9_14()
	list_9_15()
	list_9_17()
	list_9_18()
}

func calcWithTax4(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax4(price float64) float64 {
	return price
}

func list_9_4() {
	fmt.Println("======== list_9_4 ========")
	fmt.Println("=== price가 100보다 크면 20% 세금을 적용한다. ==")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		if price > 100 {
			calcFunc = calcWithTax4
		} else {
			calcFunc = calcWithoutTax4
		}
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

func calcWithTax5(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax5(price float64) float64 {
	return price
}

func list_9_5() {
	fmt.Println("======== list_9_5 ========")
	fmt.Println("===== calcFunc가 함수가 할당되었는지 체크한다. ====")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax5
		} else {
			calcFunc = calcWithoutTax5
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

func calcWithTax6(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax6(price float64) float64 {
	return price
}
func printPrice6(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func list_9_6() {
	fmt.Println("======== list_9_6 ========")
	fmt.Println("==== 인수를 함수타입으로 넘길수있다. ===")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		if price > 100 {
			printPrice6(product, price, calcWithTax6)
		} else {
			printPrice6(product, price, calcWithoutTax6)
		}
	}
}

func calcWithTax7(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax7(price float64) float64 {
	return price
}
func printPrice7(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator7(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax7
	}
	return calcWithoutTax7
}

func list_9_7() {
	fmt.Println("======== list_9_7 ========")
	fmt.Println("====== 함수의 리턴값을 함수타입으로 넘길수있다. ====")
	fmt.Println("====== selectCalculator() 리턴값이 func(float64) float64 이다. ====")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice7(product, price, selectCalculator7(price))
	}
}

type calcFunc8 func(float64) float64

func calcWithTax8(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax8(price float64) float64 {
	return price
}
func printPrice8(product string, price float64, calculator calcFunc8) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator8(price float64) calcFunc8 {
	if price > 100 {
		return calcWithTax8
	}
	return calcWithoutTax8
}

func list_9_8() {
	fmt.Println("======== list_9_8 ========")
	fmt.Println("===== type calcFunc8 func(float64) float64 을 정의하여 해당 함수타입 별칭을 사용할수있다.====")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice8(product, price, selectCalculator8(price))
	}
}

type calcFunc9 func(float64) float64

func printPrice9(product string, price float64, calculator calcFunc9) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator9(price float64) calcFunc9 {
	if price > 100 {
		var withTax calcFunc9 = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

func list_9_9() {
	fmt.Println("======== list_9_9 ========")
	fmt.Println("==== func 키워드 뒤에 함수이름을 생략하고 바로 코드블럭을 넣을수있다. ===")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice9(product, price, selectCalculator9(price))
	}
}

func list_9_10() {
	fmt.Println("======== list_9_10 ========")
	fmt.Println("==== 코드블럭 안의 변수는 코드 블럭안에서만 유효하다 ==")
}

// ----------------------------------------------------------------------------

type calcFunc11 func(float64) float64

func printPrice11(product string, price float64, calculator calcFunc11) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator11(price float64) calcFunc11 {
	if price > 100 {
		return func(price float64) float64 {
			return price + (price * 0.2)
		}
	}
	return func(price float64) float64 {
		return price
	}
}

func list_9_11() {
	fmt.Println("======== list_9_11 ========")
	fmt.Println("== 변수를 따로 할당하지 않고, 바로 리턴값으로 함수를 사용할수있다. ==")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice11(product, price, selectCalculator11(price))
	}
}

// ----------------------------------------------------------------------------

type calcFunc12 func(float64) float64

func printPrice12(product string, price float64, calculator calcFunc12) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func list_9_12() {
	fmt.Println("======== list_9_12 ========")
	fmt.Println("== 함수에 대한 인수로 익명함수를 바로 사용할수도 있다. ==")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice12(product, price, func(price float64) float64 {
			return price + (price * 0.2)
		})
	}
}

// ----------------------------------------------------------------------------

type calcFunc13 func(float64) float64

func printPrice13(product string, price float64, calculator calcFunc13) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func list_9_13() {
	fmt.Println("======== list_9_13 ========")
	fmt.Println("== 수상 스포츠와 축구에 대해서 서로 다른 세율을 계산하는 소스 ==")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	calc := func(price float64) float64 {
		if price > 100 {
			return price + (price * 0.2)
		}
		return price
	}
	for product, price := range watersportsProducts {
		printPrice13(product, price, calc)
	}
	calc = func(price float64) float64 {
		if price > 50 {
			return price + (price * 0.1)
		}
		return price
	}
	for product, price := range soccerProducts {
		printPrice13(product, price, calc)
	}
}

// ----------------------------------------------------------------------------

type calcFunc14 func(float64) float64

func printPrice14(product string, price float64, calculator calcFunc14) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func priceCalcFactory14(threshold, rate float64) calcFunc14 {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func list_9_14() {
	fmt.Println("======== list_9_14 ========")
	fmt.Println("== priceCalcFactory14 는 팩토리함수 ==")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	waterCalc := priceCalcFactory14(100, 0.2)
	soccerCalc := priceCalcFactory14(50, 0.1)
	for product, price := range watersportsProducts {
		printPrice14(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice14(product, price, soccerCalc)
	}
}

// ----------------------------------------------------------------------------

type calcFunc15 func(float64) float64

func printPrice15(product string, price float64, calculator calcFunc15) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway15 = false

func priceCalcFactory15(threshold, rate float64) calcFunc15 {
	return func(price float64) float64 {
		if prizeGiveaway15 {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func list_9_15() {
	fmt.Println("======== list_9_15 ========")
	fmt.Println("== 클로저는 함수를 호출할때 평가하기때문에, 결과값은 모두 0이 나온다. ==")
	fmt.Println("== printPrice15을 호출할당시에는 prizeGiveaway15가 0이되기때문이다.  ==")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	prizeGiveaway15 = false
	waterCalc := priceCalcFactory15(100, 0.2)
	prizeGiveaway15 = true
	soccerCalc := priceCalcFactory15(50, 0.1)
	for product, price := range watersportsProducts {
		printPrice15(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice15(product, price, soccerCalc)
	}
}

// ----------------------------------------------------------------------------

type calcFunc17 func(float64) float64

func printPrice17(product string, price float64, calculator calcFunc17) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway17 = false

func priceCalcFactory17(threshold, rate float64, zeroPrices bool) calcFunc17 {
	return func(price float64) float64 {
		if zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func list_9_17() {
	fmt.Println("======== list_9_17 ========")
	fmt.Println("== priceCalcFactory17에서 매개변수로 zeroPrices값을 받으면, 9-16번 예제의 문제점을 피할수있다. ==")

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	prizeGiveaway17 = false
	waterCalc := priceCalcFactory17(100, 0.2, prizeGiveaway17)
	prizeGiveaway17 = true
	soccerCalc := priceCalcFactory17(50, 0.1, prizeGiveaway17)
	for product, price := range watersportsProducts {
		printPrice17(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice17(product, price, soccerCalc)
	}
}

// ----------------------------------------------------------------------------

type calcFunc18 func(float64) float64

func printPrice18(product string, price float64, calculator calcFunc18) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway18 = false

func priceCalcFactory18(threshold, rate float64, zeroPrices *bool) calcFunc18 {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func list_9_18() {
	fmt.Println("======== list_9_18 ========")
	fmt.Println("== 해당 함수가 호출될때 prizeGiveaway18값이 적용되기를 원할때 포인터를 사용하여 해결할수있다. ==")
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	prizeGiveaway18 = false
	waterCalc := priceCalcFactory18(100, 0.2, &prizeGiveaway18)
	prizeGiveaway18 = true
	soccerCalc := priceCalcFactory18(50, 0.1, &prizeGiveaway18)
	for product, price := range watersportsProducts {
		printPrice18(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice18(product, price, soccerCalc)
	}
}
