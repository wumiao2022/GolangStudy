package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 1 and 10
	randomNum := rand.Intn(10) + 1

	// Print out the random number
	fmt.Println("Random number generated:", randomNum)

	// Start a loop to allow the user to guess the number
	for {
		// Ask user to guess the number
		fmt.Print("Guess the number between 1 and 10: ")
		var guess int
		fmt.Scan(&guess)

		// Check if user's guess is correct
		if guess == randomNum {
			fmt.Println("Congratulations! You guessed the number correctly!")
			break // Exit the loop if the guess is correct
		} else {
			fmt.Println("Wrong guess. Try again!")
		}
	}
}