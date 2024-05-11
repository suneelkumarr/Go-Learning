package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} else {
		fmt.Println("8 is not divisible by 4")
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		fmt.Println("myVar is Cat")
	case "Dog":
		fmt.Println("myVar is Dog")
	default:
		fmt.Println("myVar is something else")

	}
}
