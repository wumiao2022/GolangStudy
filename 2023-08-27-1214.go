package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var a int = 5
	var b float64 = 5.5
	var c string = "Golang"
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)

	// Exercise 3: Constants
	const pi float64 = 3.1415
	const studentAge int = 16
	fmt.Println("pi =", pi)
	fmt.Println("Student age =", studentAge)

	// Exercise 4: Arrays
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println("Numbers =", numbers)

	// Exercise 5: Slices
	var fruits []string = []string{"apple", "banana", "orange"}
	fmt.Println("Fruits =", fruits)

	// Exercise 6: Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Number =", i)
	}

	// Exercise 7: If-else
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Exercise 8: Functions
	result := add(3, 4)
	fmt.Println("Result =", result)
}

func add(a, b int) int {
	return a + b
}
