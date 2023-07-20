package main

import "fmt"

func main() {
	// Example 1: Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var name string = "Gopher"
	fmt.Println("Welcome,", name)

	const age int = 25
	fmt.Println("I am", age, "years old.")

	// Example 3: Arithmetic Operations
	num1 := 10
	num2 := 5
	fmt.Println("Sum:", num1+num2)
	fmt.Println("Difference:", num1-num2)
	fmt.Println("Product:", num1*num2)
	fmt.Println("Quotient:", num1/num2)

	// Example 4: Conditional Statements
	if num1 > num2 {
		fmt.Println(num1, "is greater than", num2)
	} else if num1 < num2 {
		fmt.Println(num1, "is less than", num2)
	} else {
		fmt.Println(num1, "is equal to", num2)
	}

	// Example 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Functions
	fmt.Println("Sum of 3 and 4 is:", sum(3, 4))

	// Example 7: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)
	fmt.Println("Slice:", numbers[1:4])
}

func sum(a, b int) int {
	return a + b
}