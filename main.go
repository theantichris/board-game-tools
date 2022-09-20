package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numberOfDice := flag.Int("n", 1, "The number of dice to roll. Default 1.")
	diceSides := flag.Int("d", 6, "The number of sides of the dice to roll.")
	getSum := flag.Bool("s", false, "Get the sum of the dice rolls.")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	rolls := rollDice(numberOfDice, diceSides)
	printRolls(rolls)

	if *getSum {
		sum := addRolls(rolls)

		fmt.Printf("The sum of the rolls is %d.\n", sum)
	}
}

func rollDice(numberOfDice *int, diceSides *int) []int {
	var rolls []int

	for i := 0; i < *numberOfDice; i++ {
		rolls = append(rolls, rand.Intn(*diceSides)+1)
	}

	return rolls
}

func printRolls(rolls []int) {
	for i, roll := range rolls {
		fmt.Printf("Roll %d was %d.\n", i+1, roll)
	}
}

func addRolls(rolls []int) int {
	sum := 0

	for _, roll := range rolls {
		sum += roll
	}

	return sum
}
