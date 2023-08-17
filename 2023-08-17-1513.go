package main

import "fmt"

func main() {
	// 练习1: 定义一个函数，接收两个整数参数，并返回它们的和
	func sum(a, b int) int {
		return a + b
	}

	// 练习2: 定义一个函数，接收一个整数参数n，并返回n的阶乘
	func factorial(n int) int {
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	}

	// 练习3: 定义一个函数，接收一个整数参数n，并返回斐波那契数列的第n项的值
	func fibonacci(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	// 调用上述函数并输出结果
	fmt.Println(sum(3, 5))
	fmt.Println(factorial(5))
	fmt.Println(fibonacci(7))
}