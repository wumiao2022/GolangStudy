package main

import "fmt"

func main() {
	// 1. 输出 "Hello, world!"
	fmt.Println("Hello, world!")

	// 2. 输出 1-10 的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 3. 计算 1-100 的累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否是素数
	num := 11
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 5. 使用递归实现斐波那契数列
	n := 10
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}