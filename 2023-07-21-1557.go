package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Create a variable "name" and assign it with your name, then print it
	name := "John Doe"
	fmt.Println("My name is", name)

	// Exercise 3: Create a variable "age" and assign it with your age, then print it
	age := 25
	fmt.Println("I am", age, "years old")

	// Exercise 4: Perform arithmetic operations on two variables and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num1 / num2
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", subtraction)
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)
}