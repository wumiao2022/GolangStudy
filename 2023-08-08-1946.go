package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guessNumber()
}

func guessNumber() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100) + 1

	fmt.Println("Guess a number between 1 and 100:")

	var guess int
	for {
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if guess > number {
			fmt.Println("Too high. Try again:")
		} else if guess < number {
			fmt.Println("Too low. Try again:")
		} else {
			fmt.Println("Congratulations! You guessed the correct number:", number)
			break
		}
	}
}