package main

import (
	"fmt"
)

func main() {
	// Example 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Example 2: Create a variable and print its value
	num := 3
	fmt.Println(num)

	// Example 3: Use a for loop to print out numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Example 4: Create a function that adds two numbers and return the result
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))

	// Example 5: Use a switch statement to print out a message based on the value of a variable
	num = 2
	switch num {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
	default:
		fmt.Println("The number is not 1 or 2")
	}
}