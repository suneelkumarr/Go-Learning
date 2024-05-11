package main

import (
	"fmt"
	"log"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {

	fmt.Println("Before setting myVar.FirstName")
	log.Println("Before setting myVar.FirstName")

	var myVar myStruct
	fmt.Println("After declaring myVar")
	log.Println("After declaring myVar")

	myVar.FirstName = "John"
	fmt.Println("After setting myVar.FirstName")
	log.Println("After setting myVar.FirstName")

	fmt.Println("Before creating myvar2")
	log.Println("Before creating myvar2")

	myvar2 := myStruct{FirstName: "Suneel"}
	fmt.Println("After creating myvar2")
	log.Println("After creating myvar2")

	fmt.Println(myVar.printFirstName())
	log.Println("myVar.printFirstName() returned:", myVar.printFirstName())

	fmt.Println(myvar2.printFirstName())
	log.Println("myvar2.printFirstName() returned:", myvar2.printFirstName())
}
