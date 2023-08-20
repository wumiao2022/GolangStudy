package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Add two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// Exercise 3: Print the sum of numbers from 1 to 10
	sum = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The sum of numbers from 1 to 10 is", sum)

	// Exercise 4: Calculate the factorial of a number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Printf("The factorial of %d is %d\n", num, factorial)

	// Exercise 5: Find the maximum number among three given numbers
	num1 = 10
	num2 = 20
	num3 := 15
	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}
	fmt.Println("The maximum number is", max)
}