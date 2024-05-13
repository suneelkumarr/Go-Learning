package handlers

import (
	"htmltemplatesgo/pkg/render"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
