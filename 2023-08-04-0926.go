package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var x int = 10
	var y float64 = 3.14
	var z string = "Golang"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Example 3: Arithmetic Operators
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)
	fmt.Println("Division:", x/y)
	fmt.Println("Remainder:", x%3)

	// Example 4: Conditional Statements
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// Example 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Functions
	result := addNumbers(3, 4)
	fmt.Println("Result:", result)
}

func addNumbers(a, b int) int {
	return a + b
}