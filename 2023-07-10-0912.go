package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, world!")

	// Exercise 2: Variables
	var name string = "John"
	age := 27
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// Exercise 3: Constants
	const pi = 3.14
	fmt.Println("The value of pi is", pi)

	// Exercise 4: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("The first number in the array is", numbers[0])

	// Exercise 5: Slices
	slice := numbers[1:4]
	fmt.Println("The slice is", slice)

	// Exercise 6: Loops
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// Exercise 7: Conditionals
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// Exercise 8: Functions
	sum := addNumbers(3, 5)
	fmt.Println("The sum of 3 and 5 is", sum)
}

func addNumbers(a, b int) int {
	return a + b
}