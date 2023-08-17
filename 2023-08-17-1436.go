package main

import "fmt"

func main() {
	// Exercise 1
	fmt.Println("Exercise 1:")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Exercise 2
	fmt.Println("Exercise 2:")
	x := 5
	y := 10
	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than or equal to y")
	}

	// Exercise 3
	fmt.Println("Exercise 3:")
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Print(num*2, " ")
	}
	fmt.Println()

	// Exercise 4
	fmt.Println("Exercise 4:")
	var num1 int = 10
	var num2 int = 20
	fmt.Println("Sum:", add(num1, num2))

	// Exercise 5
	fmt.Println("Exercise 5:")
	str1 := "Hello"
	str2 := "World"
	fmt.Println(concat(str1, str2))
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}

// Function to concatenate two strings
func concat(a, b string) string {
	return a + " " + b
}