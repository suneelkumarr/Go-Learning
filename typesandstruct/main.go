package main

import (
	"log"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	dateofbirth time.Time
}

func main() {
	user := User{
		firstName:   "John",
		lastName:    "Doe",
		phoneNumber: "123-456-7890",
		age:         30,
		dateofbirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user.firstName, user.dateofbirth)

}
