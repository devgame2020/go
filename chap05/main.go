package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, Operation")

	list_5_4()
	list_5_5()
	list_5_6()
	list_5_7()
	list_5_8()
	list_5_9()
	list_5_10()
	list_5_11()
	list_5_12()
	list_5_13()
	list_5_14()
	list_5_15()
	list_5_16()
	list_5_17()
	list_5_18()
	list_5_19()
	list_5_20()
	list_5_21()
	list_5_22()
	list_5_23()
	list_5_24()
	list_5_25()
	list_5_26()
	list_5_27()
	list_5_28()
	list_5_29()
	list_5_30()
	list_5_31()
	list_5_32()

}

func list_5_4() {
	fmt.Println("======== list_5_4 ========")
	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
}

func list_5_5() {
	fmt.Println("======== list_5_5 ========")
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Println(intVal * 2)                    // -2
	fmt.Println(floatVal * 2)                  // +Inf
	fmt.Println(math.IsInf((floatVal * 2), 0)) // IsInf 오버플로우 체크하는 함수
}

func list_5_6() {
	fmt.Println("======== list_5_6 ========")
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)
}

func list_5_7() {
	fmt.Println("======== list_5_7 ========")
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)
}

func list_5_8() {
	fmt.Println("======== list_5_8 ========")
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)
}

func list_5_9() {
	fmt.Println("======== list_5_9 ========")
	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
}

func list_5_10() {
	fmt.Println("======== list_5_10 ========")
	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(second == third)
	fmt.Println(second == beta)
}

func list_5_11() {
	fmt.Println("======== list_5_11 ========")
	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(*second == *third)
	fmt.Println(*second == *beta)
}

func list_5_12() {
	fmt.Println("======== list_5_12 ========")
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
}

// Error
func list_5_13() {
	fmt.Println("======== list_5_13 ========")
	// kayak := 275
	// soccerBall := 19.50
	// total := kayak + soccerBall
	// fmt.Println(total)
}

func list_5_14() {
	fmt.Println("======== list_5_14 ========")
	kayak := 275
	soccerBall := 19.50
	total := float64(kayak) + soccerBall
	fmt.Println(total)
}

func list_5_15() {
	fmt.Println("======== list_5_15 ========")
	kayak := 275
	soccerBall := 19.50
	total := kayak + int(soccerBall)
	fmt.Println(total)
	fmt.Println(int8(total))
}

func list_5_16() {
	fmt.Println("======== list_5_16 ========")
	kayak := 275
	soccerBall := 19.50
	total := kayak + int(math.Round(soccerBall))
	fmt.Println(total)
}

func list_5_17() {
	fmt.Println("======== list_5_17 ========")
	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)
}

func list_5_18() {
	fmt.Println("======== list_5_18 ========")
	val1 := "0"
	bool1, b1err := strconv.ParseBool(val1)
	if b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}

func list_5_19() {
	fmt.Println("======== list_5_19 ========")
	val1 := "0"
	if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}

func list_5_20() {
	fmt.Println("======== list_5_20 ========")
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}

func list_5_21() {
	fmt.Println("======== list_5_21 ========")
	val1 := "500"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_22() {
	fmt.Println("======== list_5_22 ========")
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_23() {
	fmt.Println("======== list_5_23 ========")
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 2, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_24() {
	fmt.Println("======== list_5_24 ========")
	val1 := "0b1100100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_25() {
	fmt.Println("======== list_5_25 ========")
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 10, 0)
	if int1err == nil {
		var intResult int = int(int1)
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_26() {
	fmt.Println("======== list_5_26 ========")
	val1 := "100"
	int1, int1err := strconv.Atoi(val1)
	if int1err == nil {
		var intResult int = int1
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func list_5_27() {
	fmt.Println("======== list_5_27 ========")
	val1 := "48.95"
	float1, float1err := strconv.ParseFloat(val1, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}
}

func list_5_28() {
	fmt.Println("======== list_5_28 ========")
	val1 := "4.895e+01"
	float1, float1err := strconv.ParseFloat(val1, 64)
	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}
}

func list_5_29() {
	fmt.Println("======== list_5_29 ========")
	val1 := true
	val2 := false
	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)
	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)
}

func list_5_30() {
	fmt.Println("======== list_5_30 ========")
	val := 275
	base10String := strconv.FormatInt(int64(val), 10)
	base2String := strconv.FormatInt(int64(val), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}

func list_5_31() {
	fmt.Println("======== list_5_31 ========")
	val := 275
	base10String := strconv.Itoa(val)
	base2String := strconv.FormatInt(int64(val), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}

func list_5_32() {
	fmt.Println("======== list_5_32 ========")
	val := 49.95
	Fstring := strconv.FormatFloat(val, 'f', 2, 64)
	Estring := strconv.FormatFloat(val, 'e', -1, 64)
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
}
