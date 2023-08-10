package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Convert a floating-point number to an integer and print the result
	floatNum := 3.14
	intNum := int(floatNum)
	fmt.Println("Integer:", intNum)

	// Exercise 4: Print the Fibonacci series up to n terms
	n := 10
	fmt.Print("Fibonacci Series: ")
	fibonacci(n)

	// Exercise 5: Reverse a string and print the result
	str := "Hello, Go!"
	reversedStr := reverseString(str)
	fmt.Println("Reversed String:", reversedStr)
}

func fibonacci(n int) {
	firstTerm := 0
	secondTerm := 1

	for i := 0; i < n; i++ {
		fmt.Print(firstTerm, " ")

		nextTerm := firstTerm + secondTerm
		firstTerm = secondTerm
		secondTerm = nextTerm
	}
	fmt.Println()
}

func reverseString(str string) string {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	return reversedStr
}
