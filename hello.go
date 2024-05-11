package main

import "fmt"

func main() {
	var username string = "World"
	fmt.Println("Hello, World!", username)
	fmt.Printf("The variable Type is : %T\n ", username)

	var isLoggedIn bool = true
	fmt.Println("Hello, World!", isLoggedIn)
	fmt.Printf("The variable Type is : %T\n ", isLoggedIn)

	var smallNumber uint8 = 255
	fmt.Println("Hello, World!", smallNumber)
	fmt.Printf("The variable Type is : %T\n ", smallNumber)

	var smallFloat float32 = 255.35563532
	fmt.Println("Hello, World!", smallFloat)
	fmt.Printf("The variable Type is : %T\n ", smallFloat)

	var anotherNumber int
	fmt.Println("Hello, World!", anotherNumber)
	fmt.Printf("The variable Type is : %T\n ", anotherNumber)

	//implicit type variable declearation
	var anotherName = "World"
	fmt.Println("Hello, World!", anotherName)
	fmt.Printf("The variable Type is : %T\n ", anotherName)

	//no var style
	anotherName1 := "World"
	fmt.Println("Hello, World!", anotherName1)
	fmt.Printf("The variable Type is : %T\n ", anotherName1)
}
