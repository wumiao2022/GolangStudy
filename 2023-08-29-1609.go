package main

import (
	"fmt"
)

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Calculate sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Print even numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Example 4: Print Fibonacci series up to n terms
	fmt.Println("Fibonacci series:")
	n := 10
	num1 = 0
	num2 = 1
	for i := 0; i < n; i++ {
		fmt.Println(num1)
		temp := num1 + num2
		num1 = num2
		num2 = temp
	}

	// Example 5: Check if a number is prime
	num := 17
	isPrime := true
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is Prime?", isPrime)
}