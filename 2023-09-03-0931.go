package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Golang Daily Practice!")

	// Exercise 1: Generate a random number between 1 and 100
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(100) + 1
	fmt.Println("Exercise 1 - Random Number:", randomNum)

	// Exercise 2: Print all even numbers between 1 and 50
	fmt.Println("Exercise 2 - Even Numbers:")
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 3: Check if a given number is prime or not
	givenNum := 17
	isPrime := true
	for i := 2; i <= givenNum/2; i++ {
		if givenNum%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Exercise 3 - Is Prime:", isPrime)

	// Exercise 4: Reverse a given string
	givenStr := "Hello, Golang!"
	reversedStr := ""
	for i := len(givenStr) - 1; i >= 0; i-- {
		reversedStr += string(givenStr[i])
	}
	fmt.Println("Exercise 4 - Reversed String:", reversedStr)
}
