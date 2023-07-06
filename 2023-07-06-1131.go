package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John Doe"
	var age int = 30
	const country string = "USA"
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)

	// Example 3: Arrays and Slices
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Array:", numbers)

	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println("Slice:", fruits)

	// Example 4: Functions
	result := add(10, 20)
	fmt.Println("Result:", result)

	// Example 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}

	// Example 6: Conditional Statements
	num := 10
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}

func add(a, b int) int {
	return a + b
}