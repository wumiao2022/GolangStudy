package main

import (
	"fmt"
)

func main() {
	// Example 1: Printing Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Calculating the summation of numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Example 3: Checking if a number is prime
	num := 29
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime:", isPrime)

	// Example 4: Reversing a string
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)

	// Example 5: Generating Fibonacci series
	n := 10
	fibSeries := []int{0, 1}
	for i := 2; i < n; i++ {
		fib := fibSeries[i-1] + fibSeries[i-2]
		fibSeries = append(fibSeries, fib)
	}
	fmt.Println("Fibonacci series:", fibSeries)
}