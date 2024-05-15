package main

import (
	"fmt"
	"htmltemplatesgo/pkg/handlers"
	"net/http"

	"htmltemplatesgo/pkg/config"
)

var PortNumber = ":8080"

//go run cmd\web\main.go

func main() {
	var app config.AppConfig
	tc, err := render.createTemplateCache()
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
