package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	age := 25
	const pi float64 = 3.14159

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("PI:", pi)

	// Example 3: Control Structures
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Example 4: Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Example 5: Functions
	sum := add(10, 5)
	fmt.Println("Sum:", sum)

	// Example 6: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	slice := numbers[1:4]
	fmt.Println("Slice:", slice)

	// Example 7: Structs
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}
	fmt.Println("Person:", person)
}

func add(a int, b int) int {
	return a + b
}