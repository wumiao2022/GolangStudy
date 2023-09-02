package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two variables, assign values to them, and print their sum
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Write a function that takes two integers as parameters and returns their product
	product := multiply(5, 8)
	fmt.Println("Product:", product)

	// Exercise 4: Create an array of strings and print each element in a new line
	arr := []string{"apple", "banana", "cherry"}
	for _, item := range arr {
		fmt.Println(item)
	}

	// Exercise 5: Write a program that determines whether a given number is even or odd
	number := 17
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func multiply(a, b int) int {
	return a * b
}