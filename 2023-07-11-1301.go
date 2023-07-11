package main

import "fmt"

func main() {
	// 1. 将两个变量的值进行交换
	a := 5
	b := 10

	// 方法一
	// tmp := a
	// a = b
	// b = tmp

	// 方法二
	a, b = b, a

	fmt.Println(a, b)

	// 2. 输出1-100之间的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 3. 使用递归实现斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}