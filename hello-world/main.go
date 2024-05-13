package main

import (
	"fmt"
	"net/http"
)

var PortNumber = ":8080"

func HomePage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))

	fmt.Fprintf(w, "This is home page")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	sum := Addvalues(10, 20)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("Sum is : %d", sum))
}

func Addvalues(x, y int) int {
	return x + y
}

func main() {
	// fmt.Println("Hello, World!")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write([]byte("Hello, World!"))
	// 	n, err := fmt.Fprint(w, "Hello, World!")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of byte written : %d", n))
	// })

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
