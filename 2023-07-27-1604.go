package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 1
	num2 := 2
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个整数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：打印出1到100之间所有的奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5：判断一个字符串是否是回文字符串
	str := "ABCCBA"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome")
	}
}