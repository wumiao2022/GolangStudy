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

	// Initialize guess as 0
	guess := 0

	fmt.Println("Guess the number between 1 and 100!")

	for guess != target {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed the correct number!")
		}
	}
}