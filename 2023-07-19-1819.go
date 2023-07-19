package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1：反转字符串
	str := "Hello, Golang!"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)

	// 练习2：判断字符串是否为回文
	isPalindrome := checkPalindrome("level")
	fmt.Println(isPalindrome)

	// 练习3：找出字符串中最长的单词
	longestWord := findLongestWord("I love programming in Golang")
	fmt.Println(longestWord)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func findLongestWord(s string) string {
	words := strings.Split(s, " ")

	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}
