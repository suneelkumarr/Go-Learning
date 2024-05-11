package main

import (
	"fmt"
)

type Animals interface {
	Says() string
	NumberOfLegs() int
}

type Bird struct {
	name         string
	numberOfLegs int
	eats         string
	livesIn      string
	canFly       bool
	canSwim      bool
	canWalk      bool
	canBreathe   bool
	canSing      bool
	canCrawl     bool
	canRun       bool
	canClimb     bool
}

func (b Bird) Says() string {
	return b.name
}

func (b Bird) NumberOfLegs() int {
	return b.numberOfLegs
}

func main() {
	bird := Bird{
		name:         "Penguin",
		numberOfLegs: 2,
		eats:         "Seeds",
		livesIn:      "Water",
		canFly:       true,
		canSwim:      false,
		canWalk:      true,
		canBreathe:   true,
		canSing:      true,
		canCrawl:     true,
		canRun:       true,
		canClimb:     true,
	}

	animals := []Animals{bird}

	for _, animal := range animals {
		fmt.Println("Name:", animal.Says(), ", Number of Legs:", animal.NumberOfLegs())
	}
}
