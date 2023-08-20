package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")
	
	// 练习2：计算并打印1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers:", sum)
	
	// 练习3：判断一个数是否是素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num % i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
	
	// 练习4：交换两个变量的值
	a := 10
	b := 20
	fmt.Println("Before swap:", a, b)
	a, b = b, a
	fmt.Println("After swap:", a, b)
	
	// 练习5：计算斐波那契数列的前n项
	n := 10
	fmt.Println("Fibonacci series:")
	fibonacci(n)
}

// 递归计算斐波那契数列
func fibonacci(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(fib(i), " ")
	}
	fmt.Println()
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}