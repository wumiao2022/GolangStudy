package main

import (
	"fmt"
	"strings"
)

func main() {
	// Exercise 1: Concatenate two strings
	str1 := "Hello,"
	str2 := " Go!"
	result := str1 + str2
	fmt.Println(result)

	// Exercise 2: Print the length and uppercase of a string
	str := "example string"
	fmt.Println("Length:", len(str))
	fmt.Println("Uppercase:", strings.ToUpper(str))

	// Exercise 3: Print the sum of numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Exercise 4: Print the Fibonacci series up to 10 elements
	a, b := 0, 1
	fmt.Print("Fibonacci Series: ")
	for i := 0; i < 10; i++ {
		fmt.Print(a, " ")
		temp := a
		a = b
		b = temp + b
	}
	fmt.Println()

	// Exercise 5: Print the factorial of a number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)
}