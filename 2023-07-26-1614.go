package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Print Name
	name := "John"
	fmt.Println("My name is", name)

	// Example 3: Sum of Two Numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 4: Looping through an Array
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example 5: Creating a Function
	result := add(5, 3)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}