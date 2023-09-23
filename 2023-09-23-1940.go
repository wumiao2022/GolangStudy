package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and initialize a variable to store your age and print it
	age := 25
	fmt.Println("My age is", age)

	// Exercise 3: Create a function to calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := calculateSum(num1, num2)
	fmt.Println("Sum:", sum)

	// Exercise 4: Create a slice with 5 elements and print each element in a separate line
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// Exercise 5: Create a map with 3 key-value pairs, where the keys are strings and the values are integers, and print each key-value pair in a separate line
	mapping := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	for key, value := range mapping {
		fmt.Println(key, ":", value)
	}
}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}