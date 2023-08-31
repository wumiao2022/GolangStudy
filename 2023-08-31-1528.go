package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, world!")

	// Example 2: Variable Declaration
	var age int = 30
	fmt.Println("I am", age, "years old.")

	// Example 3: Conditional Statement
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5.")
	} else {
		fmt.Println("Number is less than or equal to 5.")
	}

	// Example 4: Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Array
	var fruits [3]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	fmt.Println(fruits)

	// Example 6: Function
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Example 7: Package Import
	fmt.Println("Pi value is", math.Pi)
}

func add(a, b int) int {
	return a + b
}

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, world!")

	// Example 2: Variable Declaration
	var age int = 30
	fmt.Println("I am", age, "years old.")

	// Example 3: Conditional Statement
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5.")
	} else {
		fmt.Println("Number is less than or equal to 5.")
	}

	// Example 4: Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Array
	var fruits [3]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	fmt.Println(fruits)

	// Example 6: Function
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Example 7: Package Import
	fmt.Println("Pi value is", math.Pi)
}

func add(a, b int) int {
	return a + b
}