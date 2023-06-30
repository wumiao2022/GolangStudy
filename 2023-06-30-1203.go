package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 5
	num2 := 3
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Print the remainder of division
	dividend := 16
	divisor := 7
	remainder := dividend % divisor
	fmt.Println("Remainder:", remainder)

	// Exercise 4: Print the even numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 5: Print the Fibonacci sequence
	numTerms := 10
	fibonacci := []int{0, 1}
	for i := 2; i < numTerms; i++ {
		fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-1])
	}
	fmt.Println(fibonacci)
}
