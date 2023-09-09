package main

import "fmt"

func main() {
	// Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Declare a variable called "name" and assign your name to it
	name := "John Doe"
	fmt.Println("My name is", name)

	// Declare two variables called "num1" and "num2" and assign them any numbers
	num1 := 5
	num2 := 10

	// Print the sum of "num1" and "num2" to the console
	fmt.Println("The sum of", num1, "and", num2, "is", num1+num2)

	// Create a function called "addition" that takes two parameters of type int and returns their sum
	addition := func(a, b int) int {
		return a + b
	}

	// Call the "addition" function and print the result to the console
	fmt.Println("3 + 4 =", addition(3, 4))

	// Use a loop to print the numbers from 1 to 10 to the console
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Create a slice called "fruits" and add some fruit names to it
	fruits := []string{"apple", "banana", "orange"}

	// Print the length of the "fruits" slice to the console
	fmt.Println("The length of the fruits slice is", len(fruits))

	// Create a map called "grades" and add some student names and their corresponding grades to it
	grades := map[string]int{
		"John":  90,
		"Jane":  85,
		"David": 95,
	}

	// Print the grade of "John" to the console
	fmt.Println("John's grade is", grades["John"])

	// Use a switch statement to print different messages based on the value of a variable called "score"
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 80:
		fmt.Println("Good job!")
	case score >= 70:
		fmt.Println("Not bad.")
	default:
		fmt.Println("Keep practicing!")
	}
}