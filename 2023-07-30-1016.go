package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Sum two numbers and print the result.
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Calculate the area of a rectangle and print the result.
	width := 8
	height := 6
	area := width * height
	fmt.Println("Area:", area)

	// Exercise 4: Print the first 10 even numbers.
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 5: Print a multiplication table for the numbers 1 to 5.
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i*j)
		}
	}
}