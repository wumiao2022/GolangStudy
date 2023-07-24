package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Define a variable named "name" and assign it the value "John"
	name := "John"
	fmt.Println(name)

	// Exercise 3: Define a function named "add" that takes two integers as parameters and returns their sum
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 3))

	// Exercise 4: Create a slice named "numbers" with the values 1, 2, 3, 4, 5 and print it
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 5: Iterate over the "numbers" slice and print each element
	for _, num := range numbers {
		fmt.Println(num)
	}
}