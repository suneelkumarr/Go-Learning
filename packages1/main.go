package main

import (
	"log"

	"github.com/suneelkumarr/Go-Learning/tree/main/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
