package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var a int = 10
	var b float64 = 3.14
	var c string = "Golang"
	fmt.Println(a, b, c)

	// Example 3: Constants
	const x = 5
	fmt.Println(x)

	// Example 4: If statement
	num := 7
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// Example 5: For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Example 7: Functions
	result := add(3, 4)
	fmt.Println(result)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}