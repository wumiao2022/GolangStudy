package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个变量并打印它的值
	var num int = 10
	fmt.Println(num)

	// 练习3：使用 for 循环打印出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个字符串切片，包含多个不同的字符串元素，并遍历打印出来
	var names []string = []string{"Alice", "Bob", "Charlie", "Dave"}
	for _, name := range names {
		fmt.Println(name)
	}

	// 练习5：定义一个函数，接收两个整数参数，返回它们的和
	fmt.Println(add(5, 3))
}

func add(a, b int) int {
	return a + b
}