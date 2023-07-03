package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var number int
	number = 20
	const message string = "Welcome to Golang"
	fmt.Println("Number:", number)
	fmt.Println("Message:", message)

	// Example 3: Array and Slices
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	slice := array[1:3]
	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)

	// Example 4: Conditionals and Loops
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	if sum > 50 {
		fmt.Println("Sum is greater than 50")
	} else {
		fmt.Println("Sum is not greater than 50")
	}

	// Example 5: Functions
	result := calculateProduct(4, 5)
	fmt.Println("Product:", result)
}

func calculateProduct(a, b int) int {
	return a * b
}