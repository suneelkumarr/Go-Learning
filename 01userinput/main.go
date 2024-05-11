package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reder := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	// comma ok || err ok syntax

	input, _ := reder.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("numRating is ", numRating)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
	fmt.Printf("Type of this rating is %T", numRating)
}
