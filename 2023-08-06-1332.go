package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	a := 5
	b := 10
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// Example 3: Calculate the factorial of a number
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("The factorial of", n, "is", factorial)

	// Example 4: Find the maximum value in an array
	numbers := []int{4, 2, 9, 7, 5}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum value in the array", numbers, "is", max)
}