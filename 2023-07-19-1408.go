package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and initialize a string variable and print its value.
	message := "Hello, Golang!"
	fmt.Println(message)

	// Exercise 3: Declare and initialize two integer variables and print their sum.
	num1 := 5
	num2 := 3
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 4: Declare and initialize a boolean variable and print its value.
	isActive := true
	fmt.Println("Active:", isActive)

	// Exercise 5: Declare an array of strings and print each element.
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		fmt.Println(name)
	}
}