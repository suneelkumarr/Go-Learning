package main

import (
	"fmt"
	"log"
)

func main() {
	// a := 10
	// b := &a
	// fmt.Println(a, b)

	var myString string = "Hello World"
	fmt.Println(myString)
	fmt.Println(&myString)
	log.Println("My String is set to ", myString)

	changeUserPointer(&myString)
	log.Println("My String is set to ", myString)
}

func changeUserPointer(s *string) {
	newValue := "Red"
	*s = newValue

}
