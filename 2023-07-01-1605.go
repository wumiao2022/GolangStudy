package main

import "fmt"

func main() {
	// Ex1: Hello World
	fmt.Println("Hello, World!")

	// Ex2: Print a number
	fmt.Println(42)

	// Ex3: Variables
	var name string = "John"
	var age int = 25
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Ex4: Constants
	const pi float64 = 3.14159
	const daysOfWeek int = 7
	fmt.Println("PI value:", pi)
	fmt.Println("# of days in a week:", daysOfWeek)

	// Ex5: Arithmetic operations
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Ex6: Control flow - if statement
	num := 7
	if num > 0 {
		fmt.Println("Positive number")
	} else if num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}

	// Ex7: Control flow - for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Ex8: Functions
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}