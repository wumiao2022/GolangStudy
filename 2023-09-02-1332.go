package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variable Declaration and Initialization
	var age int = 20
	var name string = "John"
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Exercise 3: Basic Math Operations
	num1 := 10
	num2 := 5
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	fmt.Printf("Sum: %d, Subtraction: %d, Multiply: %d, Division: %d\n", sum, sub, mul, div)

	// Exercise 4: Using If-Else Statement
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Exercise 5: Using Switch Statement
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job!")
	case "C":
		fmt.Println("Passing grade.")
	default:
		fmt.Println("Fail.")
	}
}