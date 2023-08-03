package main

import "fmt"

func main() {
	// Exercise 1: Multiply Even Numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := multiplyEvenNumbers(numbers)
	fmt.Println(result)

	// Exercise 2: Reverse String
	str := "Hello, Golang!"
	reversed := reverseString(str)
	fmt.Println(reversed)

	// Exercise 3: Fibonacci Series
	n := 10
	fibonacciSeries := generateFibonacciSeries(n)
	fmt.Println(fibonacciSeries)
}

// Exercise 1: Multiply Even Numbers
func multiplyEvenNumbers(numbers []int) int {
	product := 1
	for _, num := range numbers {
		if num%2 == 0 {
			product *= num
		}
	}
	return product
}

// Exercise 2: Reverse String
func reverseString(str string) string {
	reversed := ""
	for _, char := range str {
		reversed = string(char) + reversed
	}
	return reversed
}

// Exercise 3: Fibonacci Series
func generateFibonacciSeries(n int) []int {
	fib := []int{}
	for i := 0; i < n; i++ {
		if i <= 1 {
			fib = append(fib, i)
		} else {
			fib = append(fib, fib[i-1]+fib[i-2])
		}
	}
	return fib
}
