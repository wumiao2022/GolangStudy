package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：计算1+2的结果并输出
	fmt.Println(1+2)

	// 练习3：输出1-100的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：求1到100之间的所有数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5：使用递归函数求N的阶乘
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}