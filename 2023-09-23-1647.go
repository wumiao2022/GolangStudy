package main

import "fmt"

func main() {
	// Exercise1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Exercise2: Print the sum of two numbers
	num1 := 5
	num2 := 3
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise3: Calculate the average of three numbers
	num3 := 10
	num4 := 15
	num5 := 20
	average := (num3 + num4 + num5) / 3
	fmt.Println("Average:", average)

	// Exercise4: Check if a number is even or odd
	num6 := 7
	if num6%2 == 0 {
		fmt.Println(num6, "is even")
	} else {
		fmt.Println(num6, "is odd")
	}

	// Exercise5: Convert temperature from Celsius to Fahrenheit
	celsius := 30
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println(celsius, "degrees Celsius is equal to", fahrenheit, "degrees Fahrenheit")
}