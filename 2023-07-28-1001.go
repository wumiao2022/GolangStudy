package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Print the first 10 numbers in the Fibonacci sequence
	fmt.Println("Fibonacci Sequence:")
	fibonacci := []int{0, 1}
	for i := 2; i < 10; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// Exercise 4: Calculate the factorial of a number and print the result
	number := 5
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", number, "is", factorial)

	// Exercise 5: Check if a number is prime and print the result
	numberToCheck := 17
	isPrime := true
	for i := 2; i <= numberToCheck/2; i++ {
		if numberToCheck%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(numberToCheck, "is prime:", isPrime)
}