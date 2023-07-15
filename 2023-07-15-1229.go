package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two integer variables and print their sum
	var a, b int = 5, 10
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// Exercise 3: Create a function that takes two integers as parameters and returns their difference
	fmt.Println(subtract(20, 8)) // Output should be 12

	// Exercise 4: Create a simple for loop that prints the numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Declare a string variable and print its length
	message := "Hello, Golang!"
	fmt.Println("Length of message:", len(message))

	// Exercise 6: Create a slice of strings and print each element
	var fruits = []string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}

func subtract(a, b int) int {
	return a - b
}