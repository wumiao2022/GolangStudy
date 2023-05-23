package main

import (
	"fmt"
)

func main() {
	// 1. 输出十个斐波那契数列
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}

	fmt.Println()

	// 2. 输出1~100中的质数
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()

	// 3. 判断字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	fmt.Printf("%s is a palindrome: %t", str, isPalindrome)
}