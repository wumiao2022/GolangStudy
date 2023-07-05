package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello World!" to the console.
	fmt.Println("Hello World!")

	// Exercise 2: Create a variable called "name" and assign your name to it. Print the value of the variable.
	name := "John Doe"
	fmt.Println("My name is", name)

	// Exercise 3: Create two variables called "num1" and "num2" and assign any two numbers to them. Print the sum of the two numbers.
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 4: Create a variable called "isEven" and assign a boolean value indicating if a number is even or not. Print the value of the variable.
	number := 7
	isEven := number%2 == 0
	fmt.Println("Is", number, "even?", isEven)

	// Exercise 5: Create an array called "fruits" and add three different fruits to it. Print the array.
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Println("Fruits:", fruits)

	// Exercise 6: Create a function called "double" that takes an integer parameter and returns its double value. Call the function and print the result.
	number := 5
	double := func(n int) int {
		return n * 2
	}
	fmt.Println("Double of", number, "is", double(number))
}