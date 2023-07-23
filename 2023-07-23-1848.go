package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Data Types
	var num1 int = 5
	var num2 float64 = 2.3
	var str string = "Golang"
	fmt.Printf("num1: %d\n", num1)
	fmt.Printf("num2: %.2f\n", num2)
	fmt.Println("str:", str)

	// Exercise 3: Basic Operators
	num3 := 10
	num4 := 3
	fmt.Println("num3 + num4 =", num3+num4)
	fmt.Println("num3 - num4 =", num3-num4)
	fmt.Println("num3 * num4 =", num3*num4)
	fmt.Println("num3 / num4 =", num3/num4)
	fmt.Println("num3 % num4 =", num3%num4)

	// Exercise 4: Conditional Statements
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Exercise 6: Arrays and Slices
	// Array
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array:", numbers)

	// Slice
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names slice:", names)

	// Exercise 7: Functions
	result := add(4, 6)
	fmt.Println("4 + 6 =", result)
}

func add(a, b int) int {
	return a + b
}