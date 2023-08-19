package main

import "fmt"

func main() {
	// Exercise: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise: Add two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise: Check if a number is even or odd
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise: Find the largest number among three numbers
	num1 = 20
	num2 = 15
	num3 := 25
	largest := num1
	if num2 > largest {
		largest = num2
	}
	if num3 > largest {
		largest = num3
	}
	fmt.Println("Largest number:", largest)

	// Exercise: Create a function to calculate the factorial of a number
	number = 5
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", number, "is", factorial)
}