package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：判断一个整数是否为奇数
	num := 7
	if num%2 == 1 {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}

	// 练习3：计算斐波那契数列的第n项
	n := 10
	fmt.Println(fibonacci(n))

	// 练习4：反转字符串
	str := "Hello, World!"
	reversed := reverseString(str)
	fmt.Println(reversed)
}

// 练习3：计算斐波那契数列的第n项
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 练习4：反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}