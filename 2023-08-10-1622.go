package main

import "fmt"

func main() {
	// 实例1：交换两个变量的值
	a := 10
	b := 20

	a, b = b, a

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 实例2：斐波那契数列
	fmt.Println(fibonacci(10))

	// 实例3：判断一个数是否是质数
	fmt.Println(isPrime(17))
}

// 斐波那契数列
func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

// 判断一个数是否是质数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}