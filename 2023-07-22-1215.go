package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, world!" to the console
	fmt.Println("Hello, world!")

	// Exercise 2: Print the sum of two numbers
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Print the average of three numbers
	num3 := 2.5
	num4 := 3.2
	num5 := 4.8
	average := (num3 + num4 + num5) / 3
	fmt.Println("The average of", num3, ",", num4, "and", num5, "is", average)

	// Exercise 4: Calculate the area of a rectangle
	width := 8.5
	height := 6.3
	area := width * height
	fmt.Println("The area of the rectangle with width", width, "and height", height, "is", area)

	// Exercise 5: Calculate the circumference of a circle
	radius := 4.2
	circumference := 2 * 3.14 * radius
	fmt.Println("The circumference of a circle with radius", radius, "is", circumference)
}