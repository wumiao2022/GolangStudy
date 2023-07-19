package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Create a variable called name and assign it with your name, then print it to the console
	name := "John Doe"
	fmt.Println(name)

	// Exercise 3: Create a slice with the numbers 1, 2, 3, 4, 5 and print each number to the console
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 4: Create a function called add that takes two integers as arguments and returns their sum
	add := func(a, b int) int {
		return a + b
	}

	// Exercise 5: Call the add function with two numbers and print the result to the console
	result := add(3, 4)
	fmt.Println(result)
}
