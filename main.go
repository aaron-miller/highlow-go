package main

import (
	"fmt"
	"math/rand"
	"time"
)

const INPUT_PROMPT_UPPER_LIMIT = "Input the upper limit: "
const INPUT_PROMPT_LOWER_LIMIT = "Input the lower limit: "
const INPUT_PROMPT_GUESS = "Guess? "

func main() {
	rand.Seed(time.Now().UnixNano())
	var limit_upper, limit_lower, guess, magic_number int
	fmt.Print(INPUT_PROMPT_UPPER_LIMIT)
	fmt.Scanf("%d", &limit_upper)

	fmt.Print(INPUT_PROMPT_LOWER_LIMIT)
	fmt.Scanf("%d", &limit_lower)

	magic_number = (rand.Intn(limit_upper-limit_lower) + limit_lower)

	for magic_number != guess {
		fmt.Print(INPUT_PROMPT_GUESS)
		fmt.Scanf("%d", &guess)

		switch {
		case guess > magic_number:
			fmt.Println("High")
		case guess < magic_number:
			fmt.Println("Low")
		case guess == magic_number:
			fmt.Println("Correct!")
		default:
			fmt.Println("Invalid input")
		}
	}
}
