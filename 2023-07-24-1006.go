package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World!")

	// 练习2: 计算和打印两个数字的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3: 打印1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习4: 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}

	// 练习5: 使用递归计算斐波那契数列的第n项
	n := 10
	fibonacci := calculateFibonacci(n)
	fmt.Printf("The %dth number in Fibonacci sequence is %d\n", n, fibonacci)
}

func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
}