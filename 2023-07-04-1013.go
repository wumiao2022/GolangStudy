package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Check if a number is even or odd and print the result
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise 4: Print all numbers from 1 to 10 using a for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a slice of fruits and print each fruit using a for-each loop
	fruits := []string{"apple", "banana", "orange", "grape"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}