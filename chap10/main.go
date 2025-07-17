package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	list_10_4()
	list_10_5()
	list_10_6()
	list_10_7()
	list_10_8()
	list_10_9()
	list_10_10()
	list_10_11()
	list_10_12()
	list_10_13()
	list_10_14()
	list_10_15()
	list_10_16()
	list_10_17()
	list_10_18()
	list_10_19()
	list_10_20()
	list_10_21()
	list_10_22()
	list_10_24()
	list_10_25()
	list_10_26()
	list_10_27()
	list_10_28()
	list_10_29()

}

func list_10_4() {
	fmt.Println("======== list_10_4 ========")

	type Product struct {
		name, category string
		price          float64
	}
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
}

func list_10_5() {
	fmt.Println("======== list_10_5 ========")
	type Product struct {
		name, category string
		price          float64
	}

	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
}

func list_10_6() {
	fmt.Println("======== list_10_6 ========")
	fmt.Println("== lifejacket은 생성만 했기때문에, 모든 필드의 값이 제로타입이다. ==")

	type Product struct {
		name, category string
		price          float64
	}
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)
}

func list_10_7() {
	fmt.Println("======== list_10_7 ========")
	fmt.Println("== 필드이름을 지정하지 않고 바로 구조체 값 설정 ==")

	type Product struct {
		name, category string
		price          float64
	}

	var kayak = Product{"Kayak", "Watersports", 275.00}
	fmt.Println("Name:", kayak.name)
	fmt.Println("Category:", kayak.category)
	fmt.Println("Price:", kayak.price)
}

func list_10_8() {
	fmt.Println("======== list_10_8 ========")
	fmt.Println("== 임베디드 필드를 사용하여, 이름을 따로 정의하지 않고 바로 사용가능하다. ==")

	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		count int
	}
	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275.00},
		count:   100,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)
}

func list_10_9() {
	fmt.Println("======== list_10_9 ========")
	fmt.Println("== 임베디드 필드는 한개만 정의가능, 두개이상 사용시 이름을 할당해야한다. ==")

	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	stockItem := StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00},
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Alt Name:", stockItem.Alternate.name)
}

func list_10_10() {
	fmt.Println("======== list_10_10 ========")
	fmt.Println("== 구조체값비교 ==")

	type Product struct {
		name, category string
		price          float64
	}
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}

func list_10_11() {
	fmt.Println("======== list_10_11 ========")
	fmt.Println("== 구조체 값 비교시 비교연산자는 슬라이스에 적용할수없다. ==")

	/*
		type Product struct {
			name, category string
			price          float64
			otherNames     []string
		}
		p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
		p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
		p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
		fmt.Println("p1 == p2:", p1 == p2)
		fmt.Println("p1 == p3:", p1 == p3)
	*/
}

func list_10_12() {
	fmt.Println("======== list_10_12 ========")
	fmt.Println("== 서로다른 구조체라도 모든필드가 동일한 이름과 타입을 갖고 있다면 비교가능하다. ==")
	type Product struct {
		name, category string
		price          float64
	}
	type Item struct {
		name     string
		category string
		price    float64
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item:", prod == Product(item))
}

func writeName13(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func list_10_13() {
	fmt.Println("======== list_10_13 ========")
	type Product struct {
		name, category string
		price          float64
		//otherNames []string
	}
	type Item struct {
		name     string
		category string
		price    float64
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName13(prod)
	writeName13(item)
}

func list_10_14() {
	fmt.Println("======== list_10_14 ========")
	type Product struct {
		name, category string
		price          float64
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())
}

func list_10_15() {
	fmt.Println("======== list_10_15 ========")
	fmt.Println("== 배열, 슬라이스 , 맵 사용예제 ==")
	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}

	array := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)
	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)

}

func list_10_16() {
	fmt.Println("======== list_10_16 ========")
	type Product struct {
		name, category string
		price          float64
	}
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 := p1
	p1.name = "Original Kayak"

	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)
}

func list_10_17() {
	fmt.Println("======== list_10_17 ========")
	type Product struct {
		name, category string
		price          float64
	}
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 := &p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", (*p2).name)
}

// =============================================================

type Product18 struct {
	name, category string
	price          float64
}

func calcTax18(product *Product18) {
	if (*product).price > 100 {
		(*product).price += (*product).price * 0.2
	}
}

func list_10_18() {
	fmt.Println("======== list_10_18 ========")
	kayak := Product18{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax18(&kayak)
	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)
}

// =============================================================

type Product19 struct {
	name, category string
	price          float64
}

func calcTax19(product *Product19) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func list_10_19() {
	fmt.Println("======== list_10_19 ========")
	kayak := Product19{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax19(&kayak)
	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)
}

// =============================================================

type Product20 struct {
	name, category string
	price          float64
}

func calcTax20(product *Product20) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func list_10_20() {
	fmt.Println("======== list_10_20 ========")
	kayak := &Product20{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax20(kayak)
	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)
}

// =============================================================

type Product21 struct {
	name, category string
	price          float64
}

func calcTax21(product *Product21) *Product21 {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func list_10_21() {
	fmt.Println("======== list_10_21 ========")
	kayak := calcTax21(&Product21{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})
	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price", kayak.price)
}

// =============================================================

type Product22 struct {
	name, category string
	price          float64
}

func newProduct22(name, category string, price float64) *Product22 {
	return &Product22{name, category, price}
}

func list_10_22() {
	fmt.Println("======== list_10_22 ========")
	fmt.Println("== 구조체 생성자함수 newProduct22() ==")

	products := [2]*Product22{
		newProduct22("Kayak", "Watersports", 275),
		newProduct22("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}
}

// =============================================================

type Product24 struct {
	name, category string
	price          float64
	*Supplier24
}
type Supplier24 struct {
	name, city string
}

func newProduct24(name, category string, price float64, supplier *Supplier24) *Product24 {
	return &Product24{name, category, price - 10, supplier}
}

func list_10_24() {
	fmt.Println("======== list_10_24 ========")
	acme := &Supplier24{"Acme Co", "New York"}
	products := [2]*Product24{
		newProduct24("Kayak", "Watersports", 275, acme),
		newProduct24("Hat", "Skiing", 42.50, acme),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier24.name, p.Supplier24.city)
	}
}

// =============================================================
type Product25 struct {
	name, category string
	price          float64
	*Supplier25
}
type Supplier25 struct {
	name, city string
}

func newProduct25(name, category string, price float64, supplier *Supplier25) *Product25 {
	return &Product25{name, category, price - 10, supplier}
}

func list_10_25() {
	fmt.Println("======== list_10_25 ========")
	acme := &Supplier25{"Acme Co", "New York"}
	p1 := newProduct25("Kayak", "Watersports", 275, acme)
	p2 := *p1
	p1.name = "Original Kayak"
	p1.Supplier25.name = "BoatCo"
	for _, p := range []Product25{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier25.name, p.Supplier25.city)
	}
}

// =============================================================

type Product26 struct {
	name, category string
	price          float64
	*Supplier26
}
type Supplier26 struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier26) *Product26 {
	return &Product26{name, category, price - 10, supplier}
}

func copyProduct(product *Product26) Product26 {
	p := *product
	s := *product.Supplier26
	p.Supplier26 = &s
	return p
}

func list_10_26() {
	fmt.Println("======== list_10_26 ========")
	acme := &Supplier26{"Acme Co", "New York"}
	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := copyProduct(p1)
	p1.name = "Original Kayak"
	p1.Supplier26.name = "BoatCo"
	for _, p := range []Product26{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier26.name, p.Supplier26.city)
	}
}

// =============================================================

type Product27 struct {
	name, category string
	price          float64
}

func list_10_27() {
	fmt.Println("======== list_10_27 ========")
	var prod Product27
	var prodPtr *Product27
	fmt.Println("Value:", prod.name, prod.category, prod.price)
	fmt.Println("Pointer:", prodPtr)
}

// =============================================================

type Product28 struct {
	name, category string
	price          float64
	*Supplier28
}
type Supplier28 struct {
	name, city string
}

func list_10_28() {
	fmt.Println("======== list_10_28 ========")
	fmt.Println("== 오류발생 : 임베디드 필드의 name에 접근해야하는데 해당 필드가 Null임 ==")
	// var prod Product28
	// var prodPtr *Product28
	// fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier28.name)
	// fmt.Println("Pointer:", prodPtr)
}

// =============================================================

func list_10_29() {
	fmt.Println("======== list_10_29 ========")
	var prod Product28 = Product28{Supplier28: &Supplier28{}}
	var prodPtr *Product28
	fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier28.name)
	fmt.Println("Pointer:", prodPtr)
}
