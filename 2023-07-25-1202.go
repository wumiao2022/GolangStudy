package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	// Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2
	// Declare a variable "name" and assign it a string value of your choice
	name := "John Doe"
	fmt.Println(name)

	// Exercise 3
	// Create a function "add" that takes two integers as parameters and returns their sum
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))

	// Exercise 4
	// Create a slice "numbers" containing integers from 1 to 5, in ascending order
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 5
	// Use a loop to print the numbers in the "numbers" slice
	for _, num := range numbers {
		fmt.Println(num)
	}
}