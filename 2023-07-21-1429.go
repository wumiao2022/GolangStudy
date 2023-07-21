package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello World!")

	// Exercise 2: Variable Declaration
	var x int = 5
	fmt.Println("The value of x is", x)

	// Exercise 3: Basic Arithmetic Operations
	y := 7
	sum := x + y
	fmt.Println("The sum of x and y is", sum)

	// Exercise 4: Conditional Statements
	if sum > 10 {
		fmt.Println("The sum is greater than 10")
	} else {
		fmt.Println("The sum is not greater than 10")
	}

	// Exercise 5: Loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("The numbers are", numbers)
}