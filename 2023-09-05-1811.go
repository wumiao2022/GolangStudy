package main

import "fmt"

func main() {
	// Exercise 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Calculate the average of three numbers
	num3 := 10
	num4 := 15
	num5 := 20
	average := (num3 + num4 + num5) / 3.0
	fmt.Println("The average of", num3, ",", num4, "and", num5, "is", average)

	// Exercise 4: Swap the values of two variables
	num6 := 4
	num7 := 6
	fmt.Println("Before swapping: num6 =", num6, "and num7 =", num7)
	temp := num6
	num6 = num7
	num7 = temp
	fmt.Println("After swapping: num6 =", num6, "and num7 =", num7)

	// Exercise 5: Print the ASCII value of a character
	char := 'A'
	fmt.Println("The ASCII value of", string(char), "is", int(char))
}