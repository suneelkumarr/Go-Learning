package main

import (
	"fmt"
	"htmltemplatesgo/pkg/config"
	"htmltemplatesgo/pkg/handlers"
	"htmltemplatesgo/pkg/render"
	"log"
	"net/http"
)

var PortNumber = ":8080"

//go run cmd\web\main.go

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
