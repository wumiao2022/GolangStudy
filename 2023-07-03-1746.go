package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Print the type and value of a variable
	myVar := "Hello, Go!"
	fmt.Printf("Type: %T, Value: %v\n", myVar, myVar)

	// Exercise 4: Use a loop to print number 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a slice of strings, append values and print the final slice
	mySlice := []string{"apple", "banana", "cherry"}
	mySlice = append(mySlice, "date", "elderberry")
	fmt.Println(mySlice)
}
