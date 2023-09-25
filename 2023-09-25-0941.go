package main

import (
	"fmt"
)

func main() {
	// 反转字符串
	str := "Hello, World!"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)

	// 求两个整数的最大公约数
	num1 := 24
	num2 := 36
	gcd := findGCD(num1, num2)
	fmt.Println(gcd)
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}