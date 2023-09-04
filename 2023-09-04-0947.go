package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Add two numbers and print the result
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("Sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Swap two numbers
	num3 := 10
	num4 := 20
	fmt.Println("Before swapping: num3 =", num3, ", num4 =", num4)
	num3, num4 = num4, num3
	fmt.Println("After swapping: num3 =", num3, ", num4 =", num4)

	// Exercise 4: Calculate the average of three numbers
	num5 := 8
	num6 := 9
	num7 := 7
	average := (num5 + num6 + num7) / 3
	fmt.Println("Average of", num5, ",", num6, "and", num7, "is", average)

	// Exercise 5: Check if a number is even or odd
	num8 := 12
	if num8%2 == 0 {
		fmt.Println(num8, "is even")
	} else {
		fmt.Println(num8, "is odd")
	}
}