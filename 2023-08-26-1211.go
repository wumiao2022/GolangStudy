package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环打印1到10的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习5：编写递归函数计算斐波那契数列的第n个数
	n := 10
	fmt.Println("Fibonacci of", n, "is", fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
