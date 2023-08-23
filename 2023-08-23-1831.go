package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算数组所有元素的和
	numbers := []int{1, 2, 3, 4, 5}
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Sum of array:", total)

	// 练习5：求斐波那契数列第n个数字
	n := 10
	fib := fibonacci(n)
	fmt.Println("Fibonacci", n, ":", fib)
}

// Fibonacci返回斐波那契数列第n个数字
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
