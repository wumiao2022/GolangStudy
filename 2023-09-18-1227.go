package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var name string = "John"
	age := 25
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Example 3: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Example 4: Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// Example 5: Functions
	result := sum(10, 20)
	fmt.Println("Sum:", result)
}

func sum(a, b int) int {
	return a + b
}