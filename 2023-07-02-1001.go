package main

import "fmt"

func main() {
	// Exercise 1: Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum of", num1, "and", num2, "is", sum)

	// Exercise 2: Create an array of integers and print each element in a new line
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// Exercise 3: Calculate the factorial of a number and print the result
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)

	// Exercise 4: Create a map of string keys and integer values and print each key-value pair
	mapping := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	for key, value := range mapping {
		fmt.Println(key, "-", value)
	}

	// Exercise 5: Check if a given string is a palindrome and print the result
	word1 := "level"
	word2 := "hello"
	isPalindrome1 := checkPalindrome(word1)
	isPalindrome2 := checkPalindrome(word2)
	fmt.Println(word1, "is a palindrome:", isPalindrome1)
	fmt.Println(word2, "is a palindrome:", isPalindrome2)
}

// Helper function to check if a string is a palindrome
func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}