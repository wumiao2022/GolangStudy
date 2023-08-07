package main

import "fmt"

func main() {
	// 练习1：实现一个计算两个数相加的函数
	add := func(a, b int) int {
		return a + b
	}

	// 练习2：实现一个计算两个数相减的函数
	subtract := func(a, b int) int {
		return a - b
	}

	// 练习3：实现一个计算两个数相乘的函数
	multiply := func(a, b int) int {
		return a * b
	}

	// 练习4：实现一个计算两个数相除的函数
	divide := func(a, b float64) float64 {
		return a / b
	}

	// 练习5：调用以上函数，输出结果
	fmt.Println("Add:", add(5, 3))
	fmt.Println("Subtract:", subtract(10, 6))
	fmt.Println("Multiply:", multiply(4, 7))
	fmt.Println("Divide:", divide(10, 2))
}