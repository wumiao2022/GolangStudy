package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Print a list of even numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Example 4: Reverse a string
	str := "Hello, World!"
	reverseStr := ""
	for _, c := range str {
		reverseStr = string(c) + reverseStr
	}
	fmt.Println("Reversed String:", reverseStr)

	// Example 5: Find the factorial of a number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)
}