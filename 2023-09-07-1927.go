package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前时间
	fmt.Println("Current time:", time.Now())

	// 练习2：计算1到100之间的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of integers from 1 to 100:", sum)

	// 练习3：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 练习4：判断一个数是否为素数
	num := 13
	isPrime := true
	for i := 2; i <= int(num/2); i++ {
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
}
