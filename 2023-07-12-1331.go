package main

import "fmt"

func main() {
	// 实现一个函数，将传入的字符串反转并返回
	str := "Hello, Go!"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)

	// 实现一个函数，计算传入的两个数的和，并返回结果
	num1 := 10
	num2 := 20
	sum := calculateSum(num1, num2)
	fmt.Println(sum)

	// 实现一个函数，判断传入的字符串是否为回文字符串，并返回结果（回文字符串是指正反排列都一样的字符串）
	word := "level"
	isPalindrome := checkPalindrome(word)
	fmt.Println(isPalindrome)
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func calculateSum(a, b int) int {
	return a + b
}

func checkPalindrome(word string) bool {
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-1-i] {
			return false
		}
	}
	return true
}
