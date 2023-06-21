package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // Use current time as seed for random number generation

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1
	fmt.Println("I am thinking of a number between 1 and 100. Can you guess it?")

	// Keep track of the number of guesses
	numGuesses := 0

	// Loop until the user guesses correctly
	for {
		// Prompt the user to enter a guess
		var guess int
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Increment the number of guesses
		numGuesses++

		// Check if the guess is too low
		if guess < target {
			fmt.Println("Too low. Try again.")
			continue
		}

		// Check if the guess is too high
		if guess > target {
			fmt.Println("Too high. Try again.")
			continue
		}

		// If we get here, the guess is correct!
		fmt.Printf("Congratulations! You guessed the number in %d tries.\n", numGuesses)
		break
	}
}