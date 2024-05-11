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

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	animals := []string{"cat", "dog", "fish", "elephant", "lion", "tiger"}

	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	newAnimal := make(map[string]string)

	newAnimal["cat"] = "tiger"
	newAnimal["dog"] = "lion"
	newAnimal["fish"] = "elephant"
	newAnimal["elephant"] = "tiger"
	newAnimal["lion"] = "tiger"
	newAnimal["tiger"] = "tiger2"

	for key, value := range newAnimal {
		fmt.Println(key, value)
	}

	var FirstLine = "This is my first Song"

	for i, l := range FirstLine {
		fmt.Println(i, l)
	}

	type Person struct {
		name        string
		age         int
		city        string
		gender      string
		nationality string
		married     bool
		height      float32
		weight      float32
		eyeColor    string
		hairColor   string
		hobbies     []string
		favFoods    []string
		about       string
		contactInfo map[string]string
	}

	var Persons []Person

	Persons = append(Persons, Person{
		name:        "Suneel",
		age:         30,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      5.11,
		weight:      60.5,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "Suneel@123",
			"phone": "1234567890",
		},
	})

	Persons = append(Persons, Person{
		name:        "kevin",
		age:         31,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      5.31,
		weight:      55.5,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "kevin@123",
			"phone": "1234567890",
		},
	})

	Persons = append(Persons, Person{
		name:        "prince",
		age:         11,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      4.11,
		weight:      40.5,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "prince@123",
			"phone": "1234567890",
		},
	})

	Persons = append(Persons, Person{
		name:        "lusios",
		age:         23,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      5.121,
		weight:      30.524,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "lusios@123",
			"phone": "1234567890",
		},
	})

	Persons = append(Persons, Person{
		name:        "gamer",
		age:         20,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      5.1,
		weight:      65.5,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "gamer@123",
			"phone": "1234567890",
		},
	})

	Persons = append(Persons, Person{
		name:        "july",
		age:         34,
		city:        "Hyderabad",
		gender:      "Male",
		nationality: "Indian",
		married:     true,
		height:      5.81,
		weight:      20.5,
		eyeColor:    "Blue",
		hairColor:   "Brown",
		hobbies:     []string{"Singing", "Dancing"},
		favFoods:    []string{"Pasta", "Chapati"},
		about:       "I am a student",
		contactInfo: map[string]string{
			"email": "july@123",
			"phone": "1234567890",
		},
	})

	for _, person := range Persons {
		fmt.Println(person)
	}

}
