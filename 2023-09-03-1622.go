package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	num1 := 5
	num2 := 8
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Iterate over an array and print each element
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:")
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example 4: Create a function that returns the Fibonacci sequence
	fmt.Println("Fibonacci Sequence:")
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}