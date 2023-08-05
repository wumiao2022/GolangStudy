package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Exercise 1: Generate a random number between 1 and 100
	number := rand.Intn(100) + 1
	fmt.Println("Random number:", number)

	// Exercise 2: Find the sum of all even numbers between 1 and 100
	sum := 0
	for i := 2; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println("Sum of even numbers:", sum)

	// Exercise 3: Print the Fibonacci sequence up to the 10th term
	first, second := 0, 1
	fmt.Print("Fibonacci sequence:", first, ", ", second)
	for i := 3; i <= 10; i++ {
		next := first + second
		fmt.Print(", ", next)
		first, second = second, next
	}
	fmt.Println()

	// Exercise 4: Print a multiplication table for numbers 1 to 10
	fmt.Println("Multiplication table:")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}