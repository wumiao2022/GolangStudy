package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Check if a number is divisible by another number
	dividend := 20
	divisor := 4
	if dividend%divisor == 0 {
		fmt.Println(dividend, "is divisible by", divisor)
	} else {
		fmt.Println(dividend, "is not divisible by", divisor)
	}

	// Exercise 4: Print all even numbers between 1 and 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 5: Reverse a string
	str := "Hello, World!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed string:", reversed)
}

// Exercise 6: Write a function to calculate the factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Exercise 7: Write a function to calculate the Fibonacci sequence up to a given number of terms
func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}