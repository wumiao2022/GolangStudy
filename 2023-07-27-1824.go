package main

import (
	"fmt"
)

func main() {
	// 练习：输出Hello World
	fmt.Println("Hello World!")

	// 练习：定义一个常量，值为10，并打印出来

	const num = 10
	fmt.Println(num)

	// 练习：定义一个变量，名为name，类型为字符串，并赋值为"Go语言"
	var name string = "Go语言"
	fmt.Println(name)

	// 练习：定义一个整型切片，并使用for循环遍历打印每个元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习：定义一个函数，名为add，接受两个整型参数并返回它们的和
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}