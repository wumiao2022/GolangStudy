package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前的日期和时间
	fmt.Println(time.Now())

	// 练习2：打印1到10的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3：计算10的阶乘
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

	// 练习4：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)
}