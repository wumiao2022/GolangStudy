package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Constants
	var x int = 5
	fmt.Println("Value of x:", x)

	const y int = 10
	fmt.Println("Value of y:", y)

	// Exercise 3: Basic Arithmetic Operations
	a := 3
	b := 4
	c := a + b
	fmt.Println("Sum of a and b:", c)

	d := a - b
	fmt.Println("Difference of a and b:", d)

	e := a * b
	fmt.Println("Product of a and b:", e)

	f := a / b
	fmt.Println("Quotient of a and b:", f)

	// Exercise 4: Conditional Statements
	num := 7
	if num < 5 {
		fmt.Println("Number is less than 5")
	} else {
		fmt.Println("Number is equal to or greater than 5")
	}

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println("Current value of i:", i)
	}

	// Exercise 6: Functions
	result := add(2, 3)
	fmt.Println("Sum of 2 and 3:", result)
}

func add(a, b int) int {
	return a + b
}