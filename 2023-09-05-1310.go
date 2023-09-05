package main

import (
	"fmt"
	"strings"
)

func main() {
	// Exercise 1: Convert string to title case
	str := "hello, how are you?"
	titleCase := strings.Title(str)
	fmt.Println(titleCase)

	// Exercise 2: Count the occurrences of each word in a string
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed faucibus nulla id feugiat imperdiet."
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	fmt.Println(wordCount)

	// Exercise 3: Reverse a string
	s := "Hello, world!"
	reverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	fmt.Println(reverse)
}