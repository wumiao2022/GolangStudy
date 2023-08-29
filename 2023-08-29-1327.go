package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variable Declaration
	var name string = "John"
	var age int = 25

	// Example 3: Basic Operations
	fmt.Println(name + " is", age, "years old.")
	fmt.Println("In 5 years, " + name + " will be", age+5, "years old.")

	// Example 4: Conditional Statements
	if age >= 18 {
		fmt.Println(name + " is an adult.")
	} else {
		fmt.Println(name + " is a minor.")
	}

	// Example 5: Loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Array
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Print(num, " ")
	}

	// Example 7: Function
	result := add(3, 5)
	fmt.Println("\nResult:", result)
}

func add(a, b int) int {
	return a + b
}
