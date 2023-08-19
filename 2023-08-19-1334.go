package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, world!")

	// Exercise 2: Variables and Constants
	num1 := 10
	num2 := 5
	result := num1 + num2
	fmt.Printf("The result of %d + %d is %d\n", num1, num2, result)

	// Exercise 3: Conditional Statements
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Exercise 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 6: Functions
	sum := add(3, 4)
	fmt.Println("The sum is:", sum)
}

func add(a, b int) int {
	return a + b
}