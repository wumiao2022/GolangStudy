package main

import "fmt"

func main() {
	// 示例1: 打印Hello, world!
	fmt.Println("Hello, world!")

	// 示例2: 变量声明和赋值
	var x int = 10
	var y int = 5
	fmt.Println("Sum:", x+y)

	// 示例3: 条件语句
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than or equal to x")
	}

	// 示例4: 循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 示例5: 函数定义和调用
	result := add(3, 4)
	fmt.Println("Result:", result)
}

// 函数: 两数相加
func add(a, b int) int {
	return a + b
}