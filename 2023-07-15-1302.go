package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	fmt.Println("Guess the number!")
	fmt.Println("Please enter a number between 1 and 100:")

	for {
		var guess int

		fmt.Scan(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100:")
			continue
		}

		if guess < target {
			fmt.Println("Too low. Guess higher:")
		} else if guess > target {
			fmt.Println("Too high. Guess lower:")
		} else {
			fmt.Println("Congratulations! You guessed the number correctly!")
			break
		}
	}
}