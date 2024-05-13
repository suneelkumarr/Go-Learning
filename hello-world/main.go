package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Hello, World!"))
		n, err := fmt.Fprint(w, "Hello, World!")

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fmt.Sprintf("Number of byte written : %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
