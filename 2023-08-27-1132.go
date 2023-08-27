package main

import (
	"fmt"
)

func main() {
	// 示例1：打印Hello World
	fmt.Println("Hello World")

	// 示例2：定义并打印一个整数变量
	var num1 int
	num1 = 10
	fmt.Println(num1)

	// 示例3：定义并打印一个字符串变量
	var str1 string
	str1 = "Go is awesome!"
	fmt.Println(str1)

	// 示例4：计算两个整数相加并打印结果
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 示例5：声明一个切片，并打印切片内容
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 示例6：使用循环遍历切片并打印每个元素
	for _, num := range nums {
		fmt.Println(num)
	}

	// 示例7：定义一个函数并调用它
	printMessage("Golang is awesome!")
}

// 示例7使用的函数，用于打印传入的消息
func printMessage(message string) {
	fmt.Println(message)
}