package main

import (
	"fmt"
)

func main() {
	// 1. 反转字符串
	str := "Hello, World!"
	reversedStr := reverseString(str)
	fmt.Println(reversedStr)

	// 2. 求两个整数的最大公约数
	x := 36
	y := 48
	gcd := findGCD(x, y)
	fmt.Println(gcd)

	// 3. 判断一个整数是否是回文数
	num := 12321
	isPalindrome := checkPalindrome(num)
	fmt.Println(isPalindrome)

	// 4. 实现斐波那契数列
	n := 10
	fibSeq := fibonacciSequence(n)
	fmt.Println(fibSeq)
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 求两个整数的最大公约数（欧几里得算法）
func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 判断一个整数是否是回文数
func checkPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	str := strconv.Itoa(n)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// 实现斐波那契数列
func fibonacciSequence(n int) []int {
	if n < 2 {
		return []int{n}
	}

	seq := make([]int, n)
	seq[0], seq[1] = 0, 1
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}