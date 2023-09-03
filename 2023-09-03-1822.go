package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Find and print the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Declare and print a string variable
	message := "Welcome to Golang!"
	fmt.Println(message)

	// Exercise 4: Use a for loop to print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a function to calculate the factorial of a number
	factorial := calculateFactorial(5)
	fmt.Println("Factorial of 5:", factorial)
}

func calculateFactorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}