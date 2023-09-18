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
	fmt.Println("I'm thinking of a number between 1 and 100.")

	// Keep track of the number of attempts
	attempts := 0

	for {
		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		attempts++

		if guess < target {
			fmt.Println("Too low! Try a higher number.")
		} else if guess > target {
			fmt.Println("Too high! Try a lower number.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}