package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Print Numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Print Even Numbers between 1 to 20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 4: Calculate the Sum of Numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 100:", sum)

	// Exercise 5: Find the Largest Number in an Array
	numbers := []int{10, 5, 8, 20, 2, 13}
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	fmt.Println("Largest number in the array:", largest)

	// Exercise 6: Generate Fibonacci Series
	n := 10
	fibonacci := generateFibonacci(n)
	fmt.Println("Fibonacci series:")
	for _, num := range fibonacci {
		fmt.Print(num, " ")
	}
}

func generateFibonacci(n int) []int {
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci
}