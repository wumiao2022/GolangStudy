package main

import (
	"fmt"
)

func main() {
	// 1. 交换两个变量的值，不使用第三个变量

	a := 10
	b := 20

	a, b = b, a

	// 2. 判断一个年份是否是闰年
	year := 2024

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}

	// 3. 判断一个字符串是否是回文字符串

	str := "level"
	isPalindrome := true

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println(str, "is a palindrome.")
	} else {
		fmt.Println(str, "is not a palindrome.")
	}

	// 4. 实现一个简单的计算器函数，接受两个数字和运算符，返回计算结果

	result := calculate(5, 3, "*")
	fmt.Println("Result:", result)

	// 5. 打印九九乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			return 0
		}
	default:
		return 0
	}
}