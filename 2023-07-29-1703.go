package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go! Let's practice programming!")

	// Example 1: Print Hello, World!
	fmt.Println("Example 1:")
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	fmt.Println("Example 2:")
	var number int = 10
	const pi float64 = 3.14159
	fmt.Println("Number:", number)
	fmt.Println("PI:", pi)

	// Example 3: If-Else statement
	fmt.Println("Example 3:")
	num1 := 5
	num2 := 10
	if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	// Example 4: For loop
	fmt.Println("Example 4:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Functions
	fmt.Println("Example 5:")
	result := add(2, 3)
	fmt.Println("The sum of 2 and 3 is:", result)

	// Example 6: Arrays and Slices
	fmt.Println("Example 6:")
	// Array
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("Numbers array:", numbers)
	// Slice
	letters := []string{"a", "b", "c"}
	fmt.Println("Letters slice:", letters)

	// Example 7: Maps
	fmt.Println("Example 7:")
	countryCapital := map[string]string{
		"US": "Washington D.C.",
		"UK": "London",
		"IN": "New Delhi",
	}
	fmt.Println("Country capital:", countryCapital)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}