package main

import (
	"fmt"
)

func main() {
	// Example 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Example 3: Use a for loop to print the numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Example 4: Use an if-else statement to check if a number is even or odd
	num := 9
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Example 5: Define a struct and create an instance of it
	type Student struct {
		Name  string
		Grade int
	}
	student := Student{Name: "John", Grade: 10}
	fmt.Println("Student:", student.Name, "Grade:", student.Grade)
}