package main

import (
	"fmt"
)

func main() {
	// 翻转字符串
	str := "Hello, World!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println(reversed)

	// 判断两个字符串是否相等
	str1 := "Hello"
	str2 := "hello"
	if str1 == str2 {
		fmt.Println("两个字符串相等")
	} else {
		fmt.Println("两个字符串不相等")
	}

	// 统计字符串中各个字符出现的次数
	str3 := "ababac"
	counts := make(map[rune]int)
	for _, char := range str3 {
		counts[char]++
	}
	fmt.Println(counts)

	// 判断一个数是否为素数
	num := 11
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}
}