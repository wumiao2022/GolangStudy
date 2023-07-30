package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Constants
	var name string = "John"
	age := 30
	const pi = 3.14

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pi:", pi)

	// Exercise 3: Arithmetic Operations
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

	// Exercise 4: Control Structures
	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is less than b")
	}

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Exercise 6: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)

	// Exercise 7: Functions
	result := add(10, 20)
	fmt.Println("Result:", result)
}

func add(a, b int) int {
	return a + b
}