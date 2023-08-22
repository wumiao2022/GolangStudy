package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var name string = "John"
	age := 30
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Example 3: Basic arithmetic operations
	num1 := 10
	num2 := 5
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Example 4: Arrays
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fmt.Println("Fruits:", fruits)

	// Example 5: Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Example 6: Functions
	result := square(5)
	fmt.Println("Square of 5:", result)
}

func square(num int) int {
	return num * num
}