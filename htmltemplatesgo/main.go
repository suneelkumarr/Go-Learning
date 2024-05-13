package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var PortNumber = ":8080"

func HomePage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	renderTemplate(w, "home.page.tmpl")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
