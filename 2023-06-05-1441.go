package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess what it is?")
	var guess int
	numAttempts := 0

	for guess != target {
		fmt.Print("Guess a number: ")
		fmt.Scan(&guess)
		numAttempts++

		if guess < target {
			fmt.Println("Too low. Guess higher.")
		} else if guess > target {
			fmt.Println("Too high. Guess lower.")
		}
	}

	fmt.Printf("You guessed it! It took you %d attempts.", numAttempts)
}