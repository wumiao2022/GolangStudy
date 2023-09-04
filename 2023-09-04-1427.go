package main

import (
	"fmt"
)

func main() {
	// Exercise 1:
	// Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2:
	// Declare a variable called "age" and assign it the value 25.
	// Print the value of "age" to the console.
	age := 25
	fmt.Println(age)

	// Exercise 3:
	// Create a function called "add" that takes in two integers and returns their sum.
	// Test the function by passing in 5 and 3, and print the result to the console.
	add := func(a, b int) int {
		return a + b
	}
	sum := add(5, 3)
	fmt.Println(sum)

	// Exercise 4:
	// Create a for loop that prints the numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5:
	// Create a slice of strings called "fruits" containing "apple", "banana", "orange".
	// Use a range loop to print each fruit on a new line.
	fruits := []string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}