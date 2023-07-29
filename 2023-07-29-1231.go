package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	fmt.Println("Hello, World!")

	// Exercise 2
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum of", num1, "and", num2, "is", sum)

	// Exercise 3
	var radius float64 = 5.0
	const pi float64 = 3.1416
	area := pi * radius * radius
	fmt.Println("Area of circle with radius", radius, "is", area)

	// Exercise 4
	var names [3]string
	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Joe"
	fmt.Println("Names:", names[0], names[1], names[2])

	// Exercise 5
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Println("Fruit", index+1, "is", fruit)
	}
}