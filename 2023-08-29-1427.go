package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print your name
	fmt.Println("Your name")

	// Exercise 3: Calculate the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 4: Determine if a number is even or odd
	number := 13
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise 5: Print the multiplication table of a number
	multiplier := 5
	for i := 1; i <= 10; i++ {
		result := multiplier * i
		fmt.Println(multiplier, "x", i, "=", result)
	}
}