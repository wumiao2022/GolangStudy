package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 打印当前时间
	now := time.Now()
	fmt.Println("Current time:", now)

	// 练习2: 判断一个数是否是素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("%d is prime: %v\n", num, isPrime)

	// 练习3: 计算斐波那契数列的第n个数
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i <= n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Printf("Fibonacci number at index %d: %d\n", n, fibonacci[n])
}