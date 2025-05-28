package main

import (
	"fmt"
	"html/template"
)

// 구조체 선언
type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// 초기사이즈 : 0
// 용량 : 10
// [] : 슬라이스
var responses = make([]*Rsvp, 0, 10)

// map을 선언하는데 key가 string이고 value가 Template
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()

	fmt.Println("TODO: add some features")
}
