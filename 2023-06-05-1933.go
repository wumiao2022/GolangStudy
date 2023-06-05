package main

import (
	"fmt"
)

func main() {
	// 练习1
	fmt.Println("Hello, world!")

	// 练习2
	fmt.Println(add(1, 2))

	// 练习3
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(s))

	// 练习4
	fmt.Println(factorial(5))

	// 练习5
	fmt.Println(fibonacci(10))
}

// 练习2：返回两个数的和
func add(a int, b int) int {
	return a + b
}

// 练习3：返回一个切片元素的总和
func sum(s []int) int {
	total := 0
	for _, value := range s {
		total += value
	}
	return total
}

// 练习4：返回一个数的阶乘
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 练习5：返回一个数列的斐波那契数列（递归实现）
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}