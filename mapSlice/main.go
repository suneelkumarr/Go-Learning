package main

import (
	"log"
)

type User struct {
	FirstName             string
	LastName              string
	Age                   int
	Email                 string
	IsActive              bool
	IsAdmin               bool
	IsSuper               bool
	IsVIP                 bool
	IsGold                bool
	IsPlatinum            bool
	IsPremium             bool
	IsVIPPlus             bool
	IsVIPPlusPlus         bool
	IsVIPPlusPlusPlus     bool
	IsVIPPlusPlusPlusPlus bool
}

func main() {
	// var myString string
	// var myInt int

	// myString = "Suneel"
	// myInt = 30

	// mySecondString := "Suneel	kumar"
	// mySecondInt := 31

	// log.Println(myString, myInt, mySecondString, mySecondInt)

	// myMap := map[string]string{
	// 	"key1": "value1",
	// 	"key2": "value2",
	// }

	// myMap := make(map[string]string)

	// myMap["dog"] = "tommy"
	// myMap["cat"] = "caty"

	// log.Println(myMap["dog"])

	// myMap := make(map[string]int)

	// myMap["dog"] = 10
	// myMap["cat"] = 20

	// log.Println(myMap["dog"])

	// myMap := make(map[string]User)

	// me := User{
	// 	FirstName:  "Suneel",
	// 	LastName:   "kumar",
	// 	Age:        31,
	// 	Email:      "suneel.kumar@go.dev",
	// 	IsActive:   true,
	// 	IsAdmin:    true,
	// 	IsSuper:    true,
	// 	IsVIP:      true,
	// 	IsGold:     true,
	// 	IsPlatinum: true,
	// 	IsPremium:  true,
	// 	IsVIPPlus:  true,
	// }

	// myMap["me"] = me

	// log.Println(myMap["me"].FirstName)

	// var mySlice []string
	// mySlice = append(mySlice, "Suneel")
	// mySlice = append(mySlice, "Kumar")
	// mySlice = append(mySlice, "Prince")
	// log.Println(mySlice)

	// var mySlice []int
	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)
	// log.Println(mySlice)

	// number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// log.Println(number)
	// log.Println(number[1:4])

	names := []string{"Suneel", "Kumar", "Prince"}
	log.Println(names)
}

// the below code fragment can be found in:
// c:/Users/techs/OneDrive/Documents/go Learning/mapSlice/main.go
