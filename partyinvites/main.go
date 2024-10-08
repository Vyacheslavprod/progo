package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend bool
}

var templates = make(map[string]*template.Template, 3)

var responses = make([]*Rsvp, 0, 10)

func loadTamplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name + ".html")
		if err != nil {
			panic(err)
		}
		templates[name] = t
		fmt.Println("Loaded template", index, name)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData {
		Rsvp: &Rsvp{}, Errors: []string {},
	})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rsvp {
		Name: request.Form["name"][0],
		Email: request.Form["email"][0],
		Phone: request.Form["phone"][0],
		WillAttend: request.Form["willattend"][0] == "true",
		}

		responses = append(responses, &responseData)

		if responseData.WillAttend {
			templates["thanks"].Execute(writer, responseData.Name)
		} else {
			templates["sorry"].Execute(writer, responseData.Name)
		}
	}
}

func main() {
	loadTamplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}