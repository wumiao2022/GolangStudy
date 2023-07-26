package main

import "fmt"

func main() {
	// 常量
	const pi = 3.14
	const message = "Hello, world!"

	// 变量
	var name string = "John"
	var age int = 25
	var isMarried bool = true

	// 数组
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// 切片
	var fruits = []string{"apple", "banana", "orange"}

	// 循环
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// 函数
	add := func(a, b int) int {
		return a + b
	}

	result := add(10, 5)
	fmt.Println(result)
}