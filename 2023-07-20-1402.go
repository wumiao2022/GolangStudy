package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	// Print current date and time
	fmt.Println("Current date and time:", time.Now())

	// Print the first 10 natural numbers
	fmt.Print("First 10 natural numbers: ")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Calculate the sum of the first 100 even numbers
	sum := 0
	for i := 2; i <= 200; i += 2 {
		sum += i
	}
	fmt.Println("Sum of the first 100 even numbers:", sum)

	// Print a triangle pattern
	size := 5
	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Calculate factorial of a number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, ":", factorial)
}