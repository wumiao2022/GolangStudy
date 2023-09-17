package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	age := 25
	const PI float64 = 3.14159

	// Example 3: Arithmetic Operators
	num1 := 10
	num2 := 5
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Example 4: If-Else Statement
	if age >= 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is a teenager.")
	}

	// Example 5: For Loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Example 6: Functions
	result := addNumbers(5, 3)
	fmt.Println("The sum of 5 and 3 is:", result)
}

func addNumbers(a, b int) int {
	return a + b
}