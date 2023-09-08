package main

import "fmt"

func main() {
	// Exercise 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Determine if a number is even or odd
	num := 27
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Exercise 4: Calculate the factorial of a number
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, "is", factorial)

	// Exercise 5: Reverse a string
	str := "Hello, Golang"
	reverse := ""
	for _, char := range str {
		reverse = string(char) + reverse
	}
	fmt.Println("Reversed string:", reverse)
}
