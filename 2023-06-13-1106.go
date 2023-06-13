package main

import (
	"fmt"
)

func main() {
	// 练习一：输出Hello World
	fmt.Println("Hello World")

	// 练习二：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 练习三：判断一个数是奇数还是偶数
	num1 := 3
	num2 := 4
	if num1%2 == 0 {
		fmt.Println(num1, "是偶数")
	} else {
		fmt.Println(num1, "是奇数")
	}
	if num2%2 == 0 {
		fmt.Println(num2, "是偶数")
	} else {
		fmt.Println(num2, "是奇数")
	}

	// 练习四：实现一个函数，判断一个字符串是否是回文串
	str1 := "level"
	str2 := "hello"
	if isPalindrome(str1) {
		fmt.Println(str1, "是回文串")
	} else {
		fmt.Println(str1, "不是回文串")
	}
	if isPalindrome(str2) {
		fmt.Println(str2, "是回文串")
	} else {
		fmt.Println(str2, "不是回文串")
	}
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}