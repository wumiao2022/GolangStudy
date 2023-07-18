package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two variables, assign values to them, and print their sum
	var num1, num2 int
	num1 = 5
	num2 = 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Create a function that takes two numbers as arguments and returns their product
	product := multiply(3, 4)
	fmt.Println("Product:", product)
}

func multiply(a, b int) int {
	return a * b
}