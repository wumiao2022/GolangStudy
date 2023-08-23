package main

import (
	"fmt"
)

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var name string = "John"
	age := 28
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Example 3: Array
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Example 4: Slice
	numbersSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numbersSlice)

	// Example 5: For loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Example 6: If-else statement
	if age > 18 {
		fmt.Println("I am an adult.")
	} else {
		fmt.Println("I am not an adult.")
	}
}