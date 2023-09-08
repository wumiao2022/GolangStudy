package main

import "fmt"

// Exercise 1: Print "Hello, World!"
func exercise1() {
	fmt.Println("Hello, World!")
}

// Exercise 2: Split a string into individual characters
func exercise2(input string) []string {
	result := make([]string, len(input))
	for i, char := range input {
		result[i] = string(char)
	}
	return result
}

// Exercise 3: Calculate the sum of two integers
func exercise3(a, b int) int {
	return a + b
}

// Exercise 4: Check if a number is even
func exercise4(num int) bool {
	return num%2 == 0
}

// Exercise 5: Reverse a string
func exercise5(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Example usage of the above exercises
func main() {
	exercise1()

	fmt.Println(exercise2("Hello"))

	fmt.Println(exercise3(3, 5))

	fmt.Println(exercise4(4))

	fmt.Println(exercise5("Hello, World!"))
}