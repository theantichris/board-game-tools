package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const dieFormatMessage = "Format: dx where x is a number. Default d6."

func main() {
	numberOfDice := flag.Int("n", 1, "The number of dice to roll. Default 1.")
	diceSides := flag.String("d", "d6", "The type of die to roll. "+dieFormatMessage)
	getSum := flag.Bool("s", false, "Get the sum of the dice rolls.")
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*diceSides))

	if matched {
		rand.Seed(time.Now().UTC().UnixNano())

		rolls := rollDice(numberOfDice, diceSides)
		printRolls(rolls)

		if *getSum {
			sum := addRolls(rolls)

			fmt.Printf("The sum of the rolls is %d.\n", sum)
		}
	} else {
		log.Fatalln("Incorrect format for die. " + dieFormatMessage)
	}
}

func rollDice(numberOfDice *int, diceSides *string) []int {
	sides := (*diceSides)[1:]
	d, err := strconv.Atoi(sides)
	if err != nil {
		log.Fatal(err)
	}

	var rolls []int

	for i := 0; i < *numberOfDice; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
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
