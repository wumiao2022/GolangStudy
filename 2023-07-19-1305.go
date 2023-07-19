package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：将两个整数相加并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：使用for循环打印1到10之间的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习4：使用切片将一个字符串反转并打印结果
	str := "Hello, Go"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println("Reversed String:", reversedStr)
}
