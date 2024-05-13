package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	myJson := `
	[
		{
			"name": "John",
			"age": 30
		},
		{
			"name": "Jane",
			"age": 25
		}
	]`

	var unmarshalData []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalData)
	if err != nil {
		log.Fatal(err)
		log.Println("Error in Unmarshal", err)
	}

	log.Printf("UnmarshalData: %v", unmarshalData)

	//write json from  struct

	var mySlice []Person
	// mySlice = append(mySlice, Person{Name: "Suneel", Age: 30})

	// jsonData, err := json.MarshalIndent(mySlice, "", " ")

	// if err != nil {
	// 	log.Fatal(err)
	// 	log.Println("Error in Marshal", err)
	// }

	// log.Println(string(jsonData))

	var m1 Person
	m1.Name = "Suneel"
	m1.Age = 30

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.Name = "Prince"
	m2.Age = 14

	mySlice = append(mySlice, m2)

	jsonData, err := json.MarshalIndent(mySlice, "", " ")

	if err != nil {
		log.Fatal(err)
		log.Println("Error in Marshal", err)
	}

	log.Println(string(jsonData))
}
