package main

import "fmt"

func main() {
	list_11_2()
	list_11_4()
	list_11_5()
	list_11_6()
	list_11_7()

	list_11_9()
	list_11_10()
	list_11_11()
	list_11_12()

	list_11_15()
	list_11_16()

	list_11_20()
	list_11_21()
	list_11_22()
	list_11_23()

	list_11_25()
	list_11_26()
	list_11_27()

}

func list_11_2() {
	fmt.Println("======== list_11_2 ========")
	fmt.Println("== 구조체를 생성하고, 구조체 내의 값들을 출력하는 예제 ==")
	type Product struct {
		name, category string
		price          float64
	}
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}
}

// ===================================================================================

type Product4 struct {
	name, category string
	price          float64
}

func printDetails4(product *Product4) {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.price)
}

func list_11_4() {
	fmt.Println("======== list_11_4 ========")
	fmt.Println("== printDetails4 함수는 Product4 구조체의 값을 출력하는 함수이다. Product4 구조체 외에 다른 곳에서는 의미없는 함수임  ==")
	products := []*Product4{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		printDetails4(p)
	}
}

// ===================================================================================
type Product5 struct {
	name, category string
	price          float64
}

func newProduct5(name, category string, price float64) *Product5 {
	return &Product5{name, category, price}
}
func (product *Product5) printDetails5() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.price)
}
func list_11_5() {
	fmt.Println("======== list_11_5 ========")
	fmt.Println("== newProduct5는 Product5 구조체를 생성하는 함수 ==")
	fmt.Println("== printDetails5는 리시버이다  ==")
	fmt.Println("== p.printDetails5() 같은 방식으로 호출한다 ==")
	products := []*Product5{
		newProduct5("Kayak", "Watersports", 275),
		newProduct5("Lifejacket", "Watersports", 48.95),
		newProduct5("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails5()
	}
}

// ===================================================================================
type Product6 struct {
	name, category string
	price          float64
}

func newProduct6(name, category string, price float64) *Product6 {
	return &Product6{name, category, price}
}
func (product *Product6) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax6(0.2, 100))
}

func (product *Product6) calcTax6(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func list_11_6() {
	fmt.Println("======== list_11_6 ========")
	fmt.Println("== calcTax6함수는 Product6의 멤버함수이고, 세금을 계산하는 함수이다. ==")
	fmt.Println("== calcTax6함수는 Product6리시버를 가지고 있고, 매개변수는 rate, threshold float64 이고, 리턴값은 float64 이다. ==")
	products := []*Product6{
		newProduct6("Kayak", "Watersports", 275),
		newProduct6("Lifejacket", "Watersports", 48.95),
		newProduct6("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails()
	}
}

// ===================================================================================
type Product7 struct {
	name, category string
	price          float64
}
type Supplier struct {
	name, city string
}

func newProduct7(name, category string, price float64) *Product7 {
	return &Product7{name, category, price}
}
func (product *Product7) printDetails7() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax7(0.2, 100))
}
func (product *Product7) calcTax7(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}
func (supplier *Supplier) printDetails7() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}

func list_11_7() {
	fmt.Println("======== list_11_7 ========")
	fmt.Println("== printDetails7 함수가 2개가 있지만 리시버가 달라서 오류가 발생하지 않는다. ==")
	fmt.Println("== 아래와 같이 같은 리시버에 같은 이름의 함수가 있으면 매개변수가 다르더라도 오류가 발생한다. ==")
	fmt.Println("== Go에서는 오버로딩을 지원하지 않는다. ==")
	fmt.Println("== func (supplier *Supplier) printDetails() ")
	fmt.Println("== func (supplier *Supplier) printDetails(showName bool) ")
	fmt.Println("====================================================================")
	products := []*Product7{
		newProduct7("Kayak", "Watersports", 275),
		newProduct7("Lifejacket", "Watersports", 48.95),
		newProduct7("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails7()
	}
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails7()
	}
}

// ===================================================================================

type Product9 struct {
	name, category string
	price          float64
}

func (product *Product9) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax9(0.2, 100))
}

func (product *Product9) calcTax9(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func list_11_9() {
	fmt.Println("======== list_11_9 ========")
	fmt.Println("== kayak변수는 Product9이지만 *Product9 리시버 함수를 호출할수있다. ==")
	fmt.Println("== ==")
	kayak := Product9{"Kayak", "Watersports", 275}
	kayak.printDetails()
}

// ===================================================================================

type Product10 struct {
	name, category string
	price          float64
}

func (product *Product10) calcTax10(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}
func (product Product10) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price", product.calcTax10(0.2, 100))
}

func list_11_10() {
	fmt.Println("======== list_11_10 ========")
	fmt.Println("== kayak변수는 *Product10 이므로 calcTax10함수호출 가능하다. ==")
	fmt.Println("== calcTax10는 리시버가 product *Product10 이지만 Product10 에서 호출가능하다 ==")
	kayak := &Product10{"Kayak", "Watersports", 275}
	kayak.printDetails()
}

// ===================================================================================

type Product11 struct {
	name, category string
	price          float64
}
type Product11List []Product11

func (products *Product11List) calcCategoryTotals11() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func list_11_11() {
	fmt.Println("======== list_11_11 ========")
	fmt.Println("== calcCategoryTotals11함수는 리시버가 products *Product11List 이다. ==")
	fmt.Println("== 리시버는 현재 패키지에 정의한 모든 타입에 대해 정의할수있다.==")
	fmt.Println("== p309 필요한 타입의 데이터를 항상수신할수있는것은 아니다. => 왜? 항상 수신하지 못하는가? ==")
	products := Product11List{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range products.calcCategoryTotals11() {
		fmt.Println("Category: ", category, "Total:", total)
	}
}

// ===================================================================================

type Product12 struct {
	name, category string
	price          float64
}
type Product12List []Product12

func (products *Product12List) calcCategoryTotals12() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}
func getProducts12() []Product12 {
	return []Product12{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

func list_11_12() {
	fmt.Println("======== list_11_12 ========")
	fmt.Println("== getProducts12()는 []Product12 를 리턴하는 함수이다. ==")
	fmt.Println("== Product12List() 는 타입변환하는 함수이다. ==")
	fmt.Println("== Product12List()는 Product12List의 리시버함수이다. ==")

	products := Product12List(getProducts12())
	for category, total := range products.calcCategoryTotals12() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	// 이건 왜 오류가 나오나?
	// calcCategoryTotals12 는 products *Product12List 에 대한 리시버이기때문에 오류
	/*
		products2 := getProducts12()
		for category, total := range products2.calcCategoryTotals12() {
			fmt.Println("Category: ", category, "Total:", total)
		}
	*/

	// 이건 왜 오류가 나오나?
	// Product12List(products3)는 값이라서, 주소를 얻지못함.
	// Product12List(products3)라는 "변환된 값"의 주소를 자동으로 얻지 못해서 *Product12List의 메서드를 호출할 수 없다.
	/*
		products3 := getProducts12()
		for category, total := range Product12List(products3).calcCategoryTotals12() {
			fmt.Println("Category: ", category, "Total:", total)
		}
	*/
}

// ===================================================================================

func list_11_15() {
	fmt.Println("======== list_11_13 ~ list_11_15 ========")
	fmt.Println("== methodsAndInterfaces15 폴더 참조 ==")
}

// ===================================================================================

func list_11_16() {
	fmt.Println("======== list_11_16 ~ list_11_19 ========")
	fmt.Println("== methodsAndInterfaces16 폴더 참조 ==")
	fmt.Println("== methodsAndInterfaces15를 interface 를 사용하여 구현하였다. ==")
	fmt.Println("== 서로 다른 struct이지만 interface를 사용하여 공통된 함수를 호출한다. ==")
}

// ===================================================================================

func list_11_20() {
	fmt.Println("======== list_11_20 ========")
	fmt.Println("== methodsAndInterfaces20 폴더 참조 ==")
	fmt.Println("== main.go파일에서 인터페이스 구현 예제 (p316) => (인터페이스 구현이 아니고 인터페이스 호출이라고 해야하지 않나?) ==")
	fmt.Println("== ==")
	fmt.Println("== ==")
}

// ===================================================================================

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

func list_11_21() {
	fmt.Println("======== list_11_21 ========")
	fmt.Println("== 구조체 필드로 인터페이스를 사용하는 예제 ==")

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))
}

// ===================================================================================

func list_11_22() {
	fmt.Println("======== list_11_22 ========")
	fmt.Println("== var expense Expense = product ==")
	fmt.Println("== product 값을 expense에 복사한다. ==")
	fmt.Println("== expense는 product의 복사본을 가지고 있다. ==")

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product
	product.price = 100
	fmt.Println(product)
	fmt.Println(expense)
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
}

// ===================================================================================

func list_11_23() {
	fmt.Println("======== list_11_23 ========")
	fmt.Println("== var expense Expense = &product 부분만 차이가 난다. ==")
	fmt.Println("== &product : 구조체값에 대한 포인터가 사용되었다. ==")

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println(product)
	fmt.Println(expense)
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))

}

// ===================================================================================

func list_11_25() {
	fmt.Println("======== list_11_25 ========")

}

// ===================================================================================

func list_11_26() {
	fmt.Println("======== list_11_26 ========")
	fmt.Println("== e1 == e2 는 포인터비교라서 false 이다. ==")
	fmt.Println("== e3 == e4 는 값비교라서 true이다. ==")

	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)
}

// ===================================================================================

func list_11_27() {
	fmt.Println("======== list_11_27 ========")
	fmt.Println("== 구조체안에 슬라이스가 있을경우 비교문에서 오류가 발생한다. ==")
	fmt.Println("== 슬라이스는 비교할수없다. ==")
}
