package main

import (
	"fmt"
)

// 练习1：判断一个数是否为素数
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 练习2：计算1~n的阶乘之和
func facSum(n int) int {
	sum := 0
	fac := 1
	for i := 1; i <= n; i++ {
		fac *= i
		sum += fac
	}
	return sum
}

// 练习3：实现一个斐波那契数列的生成器
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a + b
		a = b
		b = c
		return a
	}
}

func main() {
	// 测试isPrime函数
	fmt.Println(isPrime(7))
	fmt.Println(isPrime(15))

	// 测试facSum函数
	fmt.Println(facSum(3))
	fmt.Println(facSum(5))

	// 测试fibonacci函数
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}