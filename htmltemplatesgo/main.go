package main

import (
	"fmt"
	"net/http"
)

var PortNumber = ":8080"

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
