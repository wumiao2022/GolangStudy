package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Write a program that prints "Hello, World!" to the console.

	fmt.Println("Hello, World!")

	// Exercise 2: Write a program that calculates the sum of two numbers and prints the result.

	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Write a program that prints the first 10 even numbers.

	count := 0
	num := 0

	for count < 10 {
		if num%2 == 0 {
			fmt.Println(num)
			count++
		}
		num++
	}

	// Exercise 4: Write a program that calculates the factorial of a number and prints the result.

	num := 5
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}

	fmt.Println("Factorial of", num, "is", factorial)
}