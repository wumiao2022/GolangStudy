package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Print the product of two numbers
	num3 := 7
	num4 := 5
	product := num3 * num4
	fmt.Println("Product:", product)

	// Exercise 4: Conditionally print a message based on a boolean variable
	isEven := true
	if isEven {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// Exercise 5: Print the first 10 even numbers
	count := 0
	number := 2
	for count < 10 {
		fmt.Println(number)
		number += 2
		count++
	}
}