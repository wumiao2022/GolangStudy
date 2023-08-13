package main

import "fmt"

func main() {
	// 1. 打印 "Hello, World!" 字符串
	fmt.Println("Hello, World!")

	// 2. 定义一个整数变量 x 并初始化为 10，打印其值
	var x = 10
	fmt.Println(x)

	// 3. 定义一个字符串变量 name 并初始化为 "Golang"，打印其值
	var name = "Golang"
	fmt.Println(name)

	// 4. 计算并打印 5 + 3 的结果
	fmt.Println(5 + 3)

	// 5. 定义一个切片 s，包含 1、2、3 三个元素，打印其长度和容量
	s := []int{1, 2, 3}
	fmt.Println(len(s))
	fmt.Println(cap(s))
}