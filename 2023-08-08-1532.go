package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Constants
	const pi float64 = 3.14159
	fmt.Printf("The value of pi is %.2f.\n", pi)

	// Exercise 4: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 5: Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Exercise 6: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Conditionals
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Exercise 8: Functions
	result := add(10, 5)
	fmt.Println("Sum:", result)
}

func add(a, b int) int {
	return a + b
}