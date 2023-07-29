package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5

	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Convert a temperature from Fahrenheit to Celsius and print the result
	fahrenheit := 75.0
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)

	// Exercise 4: Combine two strings and print the result
	str1 := "Hello"
	str2 := "World"

	combined := str1 + " " + str2
	fmt.Println("Combined string:", combined)

	// Exercise 5: Find the maximum between two numbers and print the result
	num3 := 15
	num4 := 20

	max := num3
	if num4 > num3 {
		max = num4
	}
	fmt.Println("Maximum:", max)
}