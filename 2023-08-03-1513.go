package main

import "fmt"

func main() {
	// 定义一个变量并赋值为10
	num := 10

	// 使用for循环打印数字1到10
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}

	// 将字符串转换为字节数组，并打印其长度
	str := "Hello, World!"
	bytes := []byte(str)
	fmt.Println(len(bytes))

	// 定义一个切片并初始化
	names := []string{"Alice", "Bob", "Charlie"}

	// 使用for循环遍历切片并打印每个元素
	for _, name := range names {
		fmt.Println(name)
	}

	// 定义一个函数，接收两个整数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 调用函数并打印结果
	result := add(3, 5)
	fmt.Println(result)
}