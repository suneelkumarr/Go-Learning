package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := dividevales(100.0, 0)

	if err != nil {
		// w.Write([]byte(err.Error()))
		// return
		fmt.Fprintf(w, fmt.Sprintf("Error is : %s", err.Error()))
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf(" Result is : %f", f))

}

func dividevales(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
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
	http.HandleFunc("/divide", Divide)
	fmt.Printf("Starting server on port %s", "http://localhost"+PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
