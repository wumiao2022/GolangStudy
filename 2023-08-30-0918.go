package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare variables and print their values
	var a int = 10
	var b float64 = 2.5
	var c string = "Golang"
	var d bool = true

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	// Exercise 3: Perform arithmetic operations and print the result
	sum := a + int(b)
	subtraction := a - int(b)
	multiplication := a * int(b)
	division := a / int(b)

	fmt.Println("sum =", sum)
	fmt.Println("subtraction =", subtraction)
	fmt.Println("multiplication =", multiplication)
	fmt.Println("division =", division)

	// Exercise 4: Create a slice, append to it, and print its content
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers =", numbers)

	numbers = append(numbers, 6)
	fmt.Println("numbers after appending 6 =", numbers)

	// Exercise 5: Loop over a slice and print its elements
	for i := 0; i < len(numbers); i++ {
		fmt.Println("numbers[", i, "] =", numbers[i])
	}
}