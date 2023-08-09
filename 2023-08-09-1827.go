package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	age := 30
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// Exercise 3: Constants
	const pi = 3.14
	const radius = 5
	fmt.Println("The area of a circle with radius", radius, "is", pi*radius*radius)

	// Exercise 4: Data Types
	var num1 int = 10
	var num2 float64 = 2.5
	sum := num1 + int(num2)
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 5: Arrays
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println("My favorite colors are", colors[0], ",", colors[1], ", and", colors[2])

	// Exercise 6: Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("The first 3 numbers are", numbers[:3])

	// Exercise 7: Control Flow
	x := 5
	if x > 10 {
		fmt.Println(x, "is greater than 10")
	} else {
		fmt.Println(x, "is less than or equal to 10")
	}

	// Exercise 8: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println("This is loop", i)
	}

	// Exercise 9: Functions
	result := add(2, 3)
	fmt.Println("The result is", result)
}

func add(a, b int) int {
	return a + b
}