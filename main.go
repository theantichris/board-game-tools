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
	rand.Seed(time.Now().UTC().UnixNano())

	dice := flag.String("d", "d6", "The type of die to roll. "+dieFormatMessage)
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched {
		sides := (*dice)[1:]

		d, err := strconv.Atoi(sides)
		if err != nil {
			log.Fatal(err)
		}

		roll := rand.Intn(d) + 1
		fmt.Printf("You rolled a %d.\n", roll)
	} else {
		log.Fatalln("Incorrect format for die. " + dieFormatMessage)
	}
}
