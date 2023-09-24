package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go! Let's do some daily practice!")

	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Exercise 1:")
	fmt.Println("Hello, World!")

	// Exercise 2: Print the current date and time
	fmt.Println("\nExercise 2:")
	currentTime := time.Now()
	fmt.Println("Current date and time:", currentTime)

	// Exercise 3: Calculate the sum of two numbers and print the result
	fmt.Println("\nExercise 3:")
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum of", num1, "and", num2, "is", sum)

	// Exercise 4: Concatenate two strings and print the result
	fmt.Println("\nExercise 4:")
	str1 := "Hello"
	str2 := "Go!"
	result := str1 + " " + str2
	fmt.Println("Concatenated string:", result)

	// Exercise 5: Create an array of numbers and print their sum
	fmt.Println("\nExercise 5:")
	numbers := []int{1, 2, 3, 4, 5}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of array elements:", sum)
}
