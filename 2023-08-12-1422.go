package main

import "fmt"

func main() {
	// Example 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Calculate the factorial of a number
	number := 5
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", number, ":", factorial)

	// Example 4: Check if a number is even or odd
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Example 5: Iterate over an array and print its elements
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}