package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var message string = "Hello, Golang!"
	fmt.Println(message)

	// Example 3: Constants
	const pi = 3.14
	fmt.Println(pi)

	// Example 4: Data Types
	var num1 int = 10
	var num2 float64 = 3.5
	var isTrue bool = true
	fmt.Println(num1, num2, isTrue)

	// Example 5: Operators
	var a = 5
	var b = 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Example 6: Conditional Statements
	var age = 20
	if age < 18 {
		fmt.Println("You are underage.")
	} else {
		fmt.Println("You are an adult.")
	}

	// Example 7: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 8: Arrays
	var fruits [3]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	fmt.Println(fruits)

	// Example 9: Slices
	var slice1 = []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3]
	fmt.Println(slice2)

	// Example 10: Functions
	result := add(3, 4)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}