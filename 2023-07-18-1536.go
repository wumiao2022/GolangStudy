package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable "name" and assign your name to it, then print the value of "name"
	name := "Alice"
	fmt.Println(name)

	// Exercise 3: Write a program that calculates the sum of two numbers and prints the result
	a := 5
	b := 3
	sum := a + b
	fmt.Println(sum)

	// Exercise 4: Write a program that swaps the values of two variables
	x := 10
	y := 20
	x, y = y, x
	fmt.Println("x =", x, "y =", y)

	// Exercise 5: Write a program that converts Celsius to Fahrenheit
	celsius := 25.0
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println(celsius, "Celsius =", fahrenheit, "Fahrenheit")

	// Exercise 6: Write a program that checks if a number is even or odd
	num := 7
	if num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}