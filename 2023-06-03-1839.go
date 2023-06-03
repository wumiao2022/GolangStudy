package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Exercise 1: Guess the number game
	numberToGuess := rand.Intn(100)
	var guess int
	for guess != numberToGuess {
		fmt.Print("Guess the number: ")
		fmt.Scan(&guess)
		if guess < numberToGuess {
			fmt.Println("Too low!")
		} else if guess > numberToGuess {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed the number!")
		}
	}

	// Exercise 2: FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Exercise 3: Reverse string
	str := "Hello, world!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
}