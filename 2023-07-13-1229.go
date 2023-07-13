package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate and print the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Find and print the maximum number among three numbers
	num3 := 7
	if num1 > num2 && num1 > num3 {
		fmt.Println("Max:", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("Max:", num2)
	} else {
		fmt.Println("Max:", num3)
	}

	// Exercise 4: Print the first 10 Fibonacci numbers
	fibonacci := []int{0, 1}
	for i := 2; i < 10; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci:", fibonacci)

	// Exercise 5: Calculate and print the factorial of a number
	num4 := 5
	factorial := 1
	for i := 1; i <= num4; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num4, "is", factorial)
}
