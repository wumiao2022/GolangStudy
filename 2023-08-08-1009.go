package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable called "name" and assign your name to it, then print it to the console
	name := "John Doe"
	fmt.Println(name)

	// Exercise 3: Declare two variables, "x" and "y", and assign them with integer values. Add them together and print the result to the console
	x := 10
	y := 20
	sum := x + y
	fmt.Println(sum)

	// Exercise 4: Create a loop that prints numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a function called "add" that takes two integers as arguments and returns their sum
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Println(result)
}