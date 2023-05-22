package main

import (
	"fmt"
)

func main() {
	// 练习1：输出9*9乘法口诀表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习2：输入一个字符串，判断其是否为回文字符串
	var str string
	fmt.Println("请输入一个字符串：")
	fmt.Scanln(&str)
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			fmt.Println("不是回文字符串")
			return
		}
	}
	fmt.Println("是回文字符串")

	// 练习3：判断一个数是否为素数
	var num int
	fmt.Println("请输入一个正整数：")
	fmt.Scanln(&num)
	if num < 2 {
		fmt.Println("不是素数")
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Println("不是素数")
			return
		}
	}
	fmt.Println("是素数")

	// 练习4：编写一个函数，判断一个字符串是否为回文字符串
	str1 := "madams"
	fmt.Printf("%s 是否为回文字符串：%t\n", str1, isPalindrome(str1))
	str2 := "hello"
	fmt.Printf("%s 是否为回文字符串：%t\n", str2, isPalindrome(str2))
}

func isPalindrome(str string) bool {
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}