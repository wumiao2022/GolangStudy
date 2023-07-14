package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Check if a number is even or odd
	number := 17
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise 4: Reverse a string
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)

	// Exercise 5: Calculate the factorial of a number
	number = 5
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", number, "is", factorial)
}