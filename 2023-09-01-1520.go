package main

import "fmt"

func main() {
	// 示例1：打印Hello World
	fmt.Println("Hello, World!")

	// 示例2：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例3：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 示例4：计算斐波那契数列
	n := 10
	fibonacci := generateFibonacci(n)
	fmt.Println("Fibonacci Sequence:", fibonacci)
}

// 生成斐波那契数列
func generateFibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}