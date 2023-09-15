package main

import (
	"fmt"
)

func main() {
	// 1. 打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. 计算并打印 1 到 100 的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 3. 打印从 1 到 10 的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 4. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 5. 使用递归计算斐波那契数列的第 n 项
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
