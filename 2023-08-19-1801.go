package main

import "fmt"

func main() {
	// Exercise 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Print fibonacci series
	fmt.Println("Fibonacci series:")
	fibonacci(10)

	// Exercise 4: Print the first 10 odd numbers
	fmt.Println("First 10 odd numbers:")
	printOddNumbers(10)
}

func fibonacci(n int) {
	if n <= 0 {
		return
	}

	prev := 0
	curr := 1
	fmt.Print(prev, " ", curr)

	for i := 2; i < n; i++ {
		next := prev + curr
		fmt.Print(" ", next)
		prev = curr
		curr = next
	}

	fmt.Println()
}

func printOddNumbers(n int) {
	i := 1
	count := 0

	for count < n {
		if i%2 != 0 {
			fmt.Print(i, " ")
			count++
		}
		i++
	}

	fmt.Println()
}
