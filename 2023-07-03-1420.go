package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var age int = 25
	fmt.Println("My age is", age)

	// Exercise 3: If-Else
	score := 85
	if score >= 90 {
		fmt.Println("Excellent!")
	} else if score >= 80 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Keep working!")
	}

	// Exercise 4: For Loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Functions
	result := add(3, 5)
	fmt.Println("The sum is", result)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}