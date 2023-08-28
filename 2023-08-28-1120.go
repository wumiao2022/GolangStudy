package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
}

// Example 1
func example1() {
	// Write a program to print "Hello, World!" to the console.
	fmt.Println("Hello, World!")
}

// Example 2
func example2() {
	// Write a function that takes in two integers as parameters and returns their sum.
	fmt.Println(sum(3, 5)) // Output: 8
}

func sum(a, b int) int {
	return a + b
}

// Example 3
func example3() {
	// Write a program to count the number of vowels in a given string.
	str := "Hello, World!"
	count := countVowels(str)
	fmt.Println(count) // Output: 3 (count of vowels 'e', 'o', 'o')
}

func countVowels(str string) int {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	count := 0

	for _, char := range str {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
	}

	return count
}
