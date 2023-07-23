package main

import "fmt"

func main() {
	// Example 1: Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Print Fibonacci series
	n := 10
	fibonacci := []int{0, 1}

	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}

	fmt.Println(fibonacci)

	// Example 3: Calculate the factorial of a number
	num := 5
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}

	fmt.Println(factorial)

	// Example 4: Find the maximum element in an array
	arr := []int{5, 9, 3, 11, 2}
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)

	// Example 5: Check if a number is prime
	num = 13
	isPrime := true

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Println(isPrime)
}
