package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Swap the values of two variables
	a := 3
	b := 7
	a, b = b, a
	fmt.Println("a:", a, "b:", b)

	// Exercise 4: Create a function to calculate the factorial of a number
	fmt.Println("Factorial of 5:", factorial(5))

	// Exercise 5: Create a slice and append elements to it
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("Slice:", slice)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
