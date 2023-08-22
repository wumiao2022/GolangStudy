package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Calculate the factorial of a number recursively
	num := 5
	factorial := calculateFactorial(num)
	fmt.Println("Factorial of", num, ":", factorial)
}

func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}