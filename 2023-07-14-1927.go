package main

import (
	"fmt"
	"strings"
)

// Exercise 1: Reverse a string
func reverseString(s string) string {
	str := []rune(s)
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-1-i] = str[n-1-i], str[i]
	}
	return string(str)
}

// Exercise 2: Count the occurrences of a character in a string
func countOccurrences(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}

// Exercise 3: Capitalize the first letter of each word in a sentence
func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		word := []rune(words[i])
		if len(word) > 0 {
			word[0] = strings.ToUpper(string(word[0]))[0]
		}
		words[i] = string(word)
	}
	return strings.Join(words, " ")
}

func main() {
	// Test cases for the above exercises
	fmt.Println(reverseString("Hello, world!"))
	fmt.Println(countOccurrences("golang is great", 'g'))
	fmt.Println(capitalizeWords("hello, how are you?"))
}
