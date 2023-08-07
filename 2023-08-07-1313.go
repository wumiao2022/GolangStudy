package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "John"
	const age int = 30
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// Example 3: Arrays and Slices
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Example 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Functions
	result := addNumbers(3, 4)
	fmt.Println("Result:", result)

	// Example 6: Pointers
	number := 10
	changeValue(&number)
	fmt.Println("Modified Number:", number)

	// Example 7: Structs
	person := Person{name: "Alice", age: 25}
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

func addNumbers(a, b int) int {
	return a + b
}

func changeValue(num *int) {
	*num = 20
}

type Person struct {
	name string
	age  int
}