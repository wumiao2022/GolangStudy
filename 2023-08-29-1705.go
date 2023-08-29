package main

import "fmt"

func main() {
	// Exercise 1: Variable declaration and initialization
	age := 25
	fmt.Println("My age is", age)

	// Exercise 2: Basic arithmetic operations
	x := 10
	y := 5
	sum := x + y
	difference := x - y
	product := x * y
	quotient := x / y
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Exercise 3: Basic conditional statements
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}

	// Exercise 4: Looping
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Exercise 5: Functions
	result := add(x, y)
	fmt.Println("Sum:", result)
}

func add(a, b int) int {
	return a + b
}