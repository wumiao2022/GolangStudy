package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variable declaration and initialization
	var x int = 5
	fmt.Println(x)

	// Example 3: Conditional statement - if-else
	if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is non-positive")
	}

	// Example 4: Loop - for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Function declaration and calling
	result := add(3, 4)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}
