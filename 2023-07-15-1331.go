package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two variables of type int and print their sum.
	var num1, num2 int
	num1 = 10
	num2 = 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Declare a variable of type string and print its length.
	message := "This is a sample message."
	fmt.Println("Length of message:", len(message))

	// Exercise 4: Declare and initialize an array of 5 integers and print its elements.
	numbers := [5]int{2, 4, 6, 8, 10}
	fmt.Println("Numbers:", numbers)

	// Exercise 5: Use a for loop to iterate over the elements of the array and print them.
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element", i, ":", numbers[i])
	}
}