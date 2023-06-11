package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：输出1到100的数字，但是遇到3的倍数不输出，遇到5的倍数输出Buzz，遇到15的倍数输出FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%15 != 0 {
			continue
		} else if i%5 == 0 && i%15 != 0 {
			fmt.Println("Buzz")
		} else if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}

	// 练习3：定义一个函数，判断一个字符串是否是回文字符串
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("hello"))

	// 练习4：定义一个函数，将一个整数转换为二进制字符串
	fmt.Println(intToBinaryString(1001))
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func intToBinaryString(num int) string {
	var res string
	for num > 0 {
		if num%2 == 0 {
			res = "0" + res
		} else {
			res = "1" + res
		}
		num /= 2
	}
	return res
}