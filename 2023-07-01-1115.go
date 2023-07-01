package main

import "fmt"

func main() {
	// 练习1: 打印0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 声明一个字符串变量并打印出来
	str := "Hello, Golang!"
	fmt.Println(str)

	// 练习3: 声明一个整数数组并打印出来
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习4: 定义一个函数，接收两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	sum := add(2, 3)
	fmt.Println(sum)

	// 练习5: 声明一个切片，并通过切片操作获取部分元素并打印出来
	slice := []string{"apple", "banana", "cherry", "durian"}
	fmt.Println(slice[1:3])
}

请复制上方代码到Go文件中运行，并查看输出结果。