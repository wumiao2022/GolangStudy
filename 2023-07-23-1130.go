package main

import "fmt"

func main() {
	// 练习1: 将两个整数相加并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习2: 将一个字符串反转并输出结果
	str := "Hello, world!"
	reversedStr := reverseString(str)
	fmt.Println("The reversed string is:", reversedStr)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
