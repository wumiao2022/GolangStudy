package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print first 10 numbers using a for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Check if a number is even or odd
	number := 17
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 4: Calculate the sum of numbers from 1 to 100 using a while loop
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("Sum:", sum)

	// Exercise 5: Print the Fibonacci series up to a given number
	limit := 10
	fmt.Print("Fibonacci Series:", fibonacci(limit))
}

// Helper function to calculate Fibonacci series
func fibonacci(n int) []int {
	series := []int{0, 1}
	for i := 2; i < n; i++ {
		series = append(series, series[i-1]+series[i-2])
	}
	return series
}