package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	age := 30
	const pi float64 = 3.14159

	// Example 3: Arithmetic Operations
	sum := 10 + 5
	diff := 10 - 5
	product := 10 * 5
	quotient := 10 / 5

	// Example 4: Conditional Statements
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Example 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Arrays and Slices
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// Example 7: Functions
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 10))

	// Example 8: Structs
	type Person struct {
		name string
		age  int
	}
	person := Person{name: "Alice", age: 25}
	fmt.Println(person)
}