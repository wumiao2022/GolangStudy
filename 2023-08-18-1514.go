package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	var age int = 25

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 4: Conditional Statements
	num := 10

	if num < 0 {
		fmt.Println("The number is negative.")
	} else if num == 0 {
		fmt.Println("The number is zero.")
	} else {
		fmt.Println("The number is positive.")
	}

	// Exercise 5: Functions
	result := multiply(5, 3)
	fmt.Println("Multiplication result:", result)
}

func multiply(a, b int) int {
	return a * b
}
