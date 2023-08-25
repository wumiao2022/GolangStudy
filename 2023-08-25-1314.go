package main

import "fmt"

func main() {
	// 练习1：定义一个函数，实现两个整数相加的功能，并返回相加结果
	add := func(a, b int) int {
		return a + b
	}

	// 练习2：定义一个函数，实现字符串反转的功能，并返回反转后的字符串
	reverseString := func(s string) string {
		result := ""
		for i := len(s) - 1; i >= 0; i-- {
			result += string(s[i])
		}
		return result
	}

	// 练习3：定义一个函数，实现给定整数判断是否为偶数的功能，并返回判断结果
	isEven := func(num int) bool {
		return num%2 == 0
	}

	fmt.Println(add(2, 3))           // 输出：5
	fmt.Println(reverseString("hello")) // 输出：olleh
	fmt.Println(isEven(4))              // 输出：true
}