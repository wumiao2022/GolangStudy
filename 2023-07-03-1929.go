package main

import (
	"fmt"
)

func main() {
	// 练习1：将两个整数相加并打印结果
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习2：将一个字符串倒序输出
	str := "Hello, World!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed:", reversed)

	// 练习3：计算一个整数的平方和立方
	num := 10
	square := num * num
	cube := num * num * num
	fmt.Println("Square:", square)
	fmt.Println("Cube:", cube)
}
