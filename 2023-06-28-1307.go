package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable and assign it a value
	var num int = 5

	// Exercise 3: Perform arithmetic operations on the variable
	num += 10           // Addition
	num -= 5            // Subtraction
	num *= 2            // Multiplication
	num /= 3            // Division
	num %= 4            // Modulus
	num++               // Increment
	num--               // Decrement

	// Exercise 4: Use a loop to print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Define a function that calculates the sum of two numbers and returns the result
	sum := add(3, 4)
	fmt.Println(sum)

	// Exercise 6: Define a struct type and create an instance of it
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John", Age: 30}
	fmt.Println(person)
}

// Function for Exercise 5
func add(a, b int) int {
	return a + b
}
