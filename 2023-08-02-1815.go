package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two integers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// Exercise 3: Print the even numbers between 1 and 10
	fmt.Println("Even numbers between 1 and 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 4: Create a function to calculate the factorial of a given number
	fmt.Println("Factorial of 5:", factorial(5))
}

// Function to calculate the factorial of a given number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}