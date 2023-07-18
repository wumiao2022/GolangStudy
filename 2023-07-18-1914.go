package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：交换两个变量的值
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)

	// 练习3：计算斐波那契数列的第n个数
	n := 10
	fmt.Println(fibonacci(n))

	// 练习4：判断一个数是否是素数
	num := 37
	if isPrime(num) {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

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