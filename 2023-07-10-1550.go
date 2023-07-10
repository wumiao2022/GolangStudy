package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：定义一个函数，计算两个整数的和并返回结果
	sum := add(3, 5)
	fmt.Println("Sum:", sum)
}

func add(a, b int) int {
	return a + b
}