package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	a := 5
	b := 7
	sum := a + b
	fmt.Println("Sum:", sum)

	// Exercise 3: Reverse a string
	str := "hello"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)

	// Exercise 4: Calculate the factorial of a number
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, ":", factorial)

	// Exercise 5: Calculate the Fibonacci sequence
	terms := 10
	firstTerm := 0
	secondTerm := 1
	fibonacci := []int{firstTerm, secondTerm}
	for i := 2; i < terms; i++ {
		nextTerm := firstTerm + secondTerm
		fibonacci = append(fibonacci, nextTerm)
		firstTerm = secondTerm
		secondTerm = nextTerm
	}
	fmt.Println("Fibonacci Sequence:", fibonacci)
}
