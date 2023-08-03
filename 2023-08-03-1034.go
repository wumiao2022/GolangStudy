package main

import "fmt"

func main() {
	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 计算一个整数的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(n, "factorial is", factorial)

	// 判断一个字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "is a palindrome")
	} else {
		fmt.Println(str, "is not a palindrome")
	}
}

// 输出：
// Hello, World!
// Sum: 30
// 15 is odd
// 5 factorial is 120
// level is a palindrome