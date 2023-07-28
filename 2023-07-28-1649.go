package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// Exercise 3: Calculate the average of three numbers
	num1 := 12
	num2 := 17
	num3 := 20
	average := (num1 + num2 + num3) / 3
	fmt.Println("Average:", average)

	// Exercise 4: Create a for loop to print numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Print a pattern of asterisks
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Exercise 6: Calculate the factorial of a number
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, "is", factorial)
}