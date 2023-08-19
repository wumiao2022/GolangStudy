package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成一个随机数，判断该数是偶数还是奇数

	num := rand.Intn(100) // 生成0-100之间的随机数
	
	if num%2 == 0 {
		fmt.Printf("%d 是偶数。\n", num)
	} else {
		fmt.Printf("%d 是奇数。\n", num)
	}

	// 练习2：判断一个字符串是否为回文字符串

	str1 := "abcba" // 回文字符串
	str2 := "hello" // 非回文字符串

	isPalindrome1 := true
	for i := 0; i < len(str1)/2; i++ {
		if str1[i] != str1[len(str1)-1-i] {
			isPalindrome1 = false
			break
		}
	}

	if isPalindrome1 {
		fmt.Printf("%s 是回文字符串。\n", str1)
	} else {
		fmt.Printf("%s 不是回文字符串。\n", str1)
	}

	isPalindrome2 := true
	for i := 0; i < len(str2)/2; i++ {
		if str2[i] != str2[len(str2)-1-i] {
			isPalindrome2 = false
			break
		}
	}

	if isPalindrome2 {
		fmt.Printf("%s 是回文字符串。\n", str2)
	} else {
		fmt.Printf("%s 不是回文字符串。\n", str2)
	}

	// 练习3：计算一个整数的阶乘

	num2 := 5 // 要计算阶乘的数

	factorial := 1
	for i := 1; i <= num2; i++ {
		factorial *= i
	}

	fmt.Printf("%d 的阶乘为 %d。\n", num2, factorial)
}