package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	age := 25
	const PI float64 = 3.14

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("PI: %f\n\n", PI)

	// Example 3: Arithmetic Operations
	num1 := 10
	num2 := 5

	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	remainder := num1 % num2

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
	fmt.Printf("Remainder: %d\n\n", remainder)

	// Example 4: Conditional Statements
	age = 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
}