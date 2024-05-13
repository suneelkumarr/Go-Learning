package main

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	renderTemplate(w, "home.page.tmpl")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}
