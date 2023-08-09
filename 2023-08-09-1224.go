package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	const age int = 30
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")

	// Example 3: Basic Arithmetic Operations
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Example 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// Example 5: Conditional Statements
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Example 6: Arrays
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)

	// Example 7: Functions
	result := add(5, 3)
	fmt.Println("Sum:", result)
}

func add(a, b int) int {
	return a + b
}
