package main

import (
	"fmt"
)

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variable declaration and initialization
	var age int = 25
	fmt.Println("My age is", age)

	// Example 3: Simple if else statement
	num := 10
	if num > 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negative")
	}

	// Example 4: For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Function declaration and call
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}