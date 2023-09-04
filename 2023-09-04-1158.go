package main

import (
	"fmt"
)

func main() {
	// 1. 将字符串 "Hello, World!" 输出到控制台
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并初始化为10，将其打印到控制台
	num := 10
	fmt.Println(num)

	// 3. 创建一个字符串切片，并打印切片的长度和容量
	strSlice := []string{"apple", "banana", "orange"}
	fmt.Println("Length:", len(strSlice))
	fmt.Println("Capacity:", cap(strSlice))

	// 4. 使用循环打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 5. 定义一个函数，接收两个整数，返回它们的和
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}
