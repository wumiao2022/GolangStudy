package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two variables, a and b, with initial values of 5 and 3 respectively.
	a := 5
	b := 3

	// Exercise 3: Add a and b and store the result in a new variable called sum.
	sum := a + b
	fmt.Println("Sum:", sum)

	// Exercise 4: Create a slice of integers containing the numbers 1, 2, and 3.
	numbers := []int{1, 2, 3}

	// Exercise 5: Iterate through the slice and print each number.
	for _, num := range numbers {
		fmt.Println(num)
	}
}