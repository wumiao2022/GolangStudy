package main

import "fmt"

func main() {
	// 示例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 示例2：变量声明和赋值
	var num1 int = 10
	var num2 int = 20
	fmt.Println("Sum:", num1+num2)

	// 示例3：条件判断和循环
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 示例4：函数定义和调用
	result := add(5, 3)
	fmt.Println("Result:", result)
}

// 函数：计算两个数字的和
func add(a, b int) int {
	return a + b
}