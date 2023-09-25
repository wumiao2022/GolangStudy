package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two integers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Create a loop that prints numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 4: Declare a function that takes two integers as parameters and returns their product
	fmt.Println("Product:", multiply(4, 5))

	// Exercise 5: Create a slice of strings and print each element
	strings := []string{"apple", "banana", "orange"}
	for _, str := range strings {
		fmt.Println(str)
	}
}

func multiply(num1, num2 int) int {
	return num1 * num2
}