package main

import (
	"fmt"
)

// 练习1：输出九九乘法表
func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Println()
	}
}

// 练习2：计算斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 练习3：判断一个数是否是素数
func isPrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 练习4：反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 练习5：判断字符串是否是回文字符串
func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// 调用练习1
	multiplicationTable()

	// 调用练习2
	fmt.Printf("第10个斐波那契数列值为：%d\n", fibonacci(10))

	// 调用练习3
	fmt.Println("判断17是否是素数：", isPrimeNumber(17))

	// 调用练习4
	fmt.Println("反转字符串 \"Hello, World!\"：", reverseString("Hello, World!"))

	// 调用练习5
	fmt.Println("判断字符串 \"level\" 是否是回文字符串：", isPalindrome("level"))
}