package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!" to console
	fmt.Println("Hello, World!")

	// Example 2: Fibonacci sequence generator
	fibonacci(10)

	// Example 3: Calculate factorial of a number
	fmt.Println(factorial(5))
}

// Example 2: Fibonacci sequence generator
func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

// Example 3: Calculate factorial of a number
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}