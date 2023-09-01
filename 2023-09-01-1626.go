package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello, Golang!")
	fmt.Println("Today's exercise challenge is to generate a random number between 1 and 10.")

	randomNumber := generateRandomNumber(1, 10)
	fmt.Printf("The random number is: %d", randomNumber)
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}