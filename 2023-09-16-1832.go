package main

import "fmt"

func main() {
	// Example 1: Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Sum of Two Numbers
	num1 := 10
	num2 := 15
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Check Even or Odd
	num := 22
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Example 4: Reverse a String
	str := "Hello, World!"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println("Reversed String:", reverseStr)

	// Example 5: Fibonacci Series
	numTerms := 10
	fibonacciSeries := []int{0, 1}
	for i := 2; i < numTerms; i++ {
		fibonacciSeries = append(fibonacciSeries, fibonacciSeries[i-1]+fibonacciSeries[i-2])
	}
	fmt.Println("Fibonacci Series:", fibonacciSeries)
}