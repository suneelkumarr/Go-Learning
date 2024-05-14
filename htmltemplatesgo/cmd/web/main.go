package main

import (
	"fmt"
	"htmltemplatesgo/pkg/handlers"
	"net/http"
)

var PortNumber = ":8080"

//go run cmd\web\main.go

func main() {

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
