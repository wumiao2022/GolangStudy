package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	age := 25
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// Exercise 3: Constants
	const pi = 3.14159
	fmt.Println("The value of pi is", pi)

	// Exercise 4: Arithmetic operations
	x := 10
	y := 5
	sum := x + y
	difference := x - y
	product := x * y
	quotient := x / y
	fmt.Println("Sum:", sum, "Difference:", difference, "Product:", product, "Quotient:", quotient)

	// Exercise 5: If-else statements
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Exercise 6: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 8: Slices
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	// Exercise 9: Maps
	ages := map[string]int{
		"John":  25,
		"Alice": 30,
		"Bob":   35,
	}
	fmt.Println(ages)

	// Exercise 10: Functions
	result := addNumbers(10, 20)
	fmt.Println("Sum of numbers:", result)
}

func addNumbers(a, b int) int {
	return a + b
}