package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算并打印1+2的结果
	sum := 1 + 2
	fmt.Println("1 + 2 =", sum)

	// 练习3：实现一个函数，接收两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("5 + 3 =", add(5, 3))

	// 练习4：实现一个函数，接收一个字符串并返回其长度
	getLength := func(str string) int {
		return len(str)
	}
	fmt.Println("字符串\"Golang\"的长度为", getLength("Golang"))
}