package main

import "fmt"

func main() {
	// 1. 输出 Hello World
	fmt.Println("Hello World")

	// 2. 求两个数的和并输出结果
	var a, b int = 5, 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 3. 判断一个数的正负并输出结果
	num := -5
	if num >= 0 {
		fmt.Println("Positive number")
	} else {
		fmt.Println("Negative number")
	}

	// 4. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 5. 使用递归实现阶乘并输出结果
	n := 5
	fact := factorial(n)
	fmt.Println("Factorial:", fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}