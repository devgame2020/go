package main

import (
	"fmt"
	"html/template"
	"net/http"
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

// map을 선언하는데 key가 string이고 value가 *template.Template
// 초기 용량은 3
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		// layout.html의 기본형식에 각각의 html이 삽입된다.
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// --------------------- http --------------------------------------------------
func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	for i, user := range responses {
		fmt.Printf("User %d:\n", i+1)
		fmt.Printf("  Name: %s\n", user.Name)
		fmt.Printf("  Email: %s\n", user.Email)
		fmt.Printf("  Phone: %s\n", user.Phone)
		fmt.Printf("  Will Attend: %t\n\n", user.WillAttend)
	}
	templates["list"].Execute(writer, responses)
}

// ----------------------------------------------------------------------

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		// Rsvp구조체에 name,email등등의 form에서 입력한 값이 저장된다.
		responseData := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}

		// 각각 입력값에 대한 오류처리를 한다.
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			fmt.Println("responses append", responseData.Name)
			// 오류가 없으면 responses에 append한다.
			responses = append(responses, &responseData)
			if responseData.WillAttend {
				templates["thanks"].Execute(writer, responseData.Name)
			} else {
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}
	}
}

func main() {
	loadTemplates()

	//
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	// 5000번 Port로 대기한다.
	// 5000번 port로 http get,post요청이 오면 HandleFunc()에서 맵핑된 함수가 호출된다.
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
