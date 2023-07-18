package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Constants
	var name string = "John"
	age := 25
	const pi float64 = 3.14

	// Exercise 3: Print Variables
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
	fmt.Println("The value of pi is", pi)

	// Exercise 4: Basic Math Operations
	a := 10
	b := 5

	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Exercise 5: Conditional Statements
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Exercise 6: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Arrays and Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 8: Functions
	result := add(3, 4)
	fmt.Println("The sum is", result)
}

func add(a int, b int) int {
	return a + b
}