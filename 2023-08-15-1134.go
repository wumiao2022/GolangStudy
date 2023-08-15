package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// Exercise 3: Print the average of three numbers
	num3 := 8
	average := (num1 + num2 + num3) / 3
	fmt.Println("The average is:", average)

	// Exercise 4: Print the square root of a number
	num4 := 16.0
	sqrt := math.Sqrt(num4)
	fmt.Println("The square root is:", sqrt)

	// Exercise 5: Print the factorial of a number
	num5 := 5
	factorial := 1
	for i := 1; i <= num5; i++ {
		factorial *= i
	}
	fmt.Println("The factorial is:", factorial)
}