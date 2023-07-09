package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前的日期和时间
	fmt.Println(time.Now())

	// 练习2：计算1到100的和，并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= int(num/2); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(num, "is prime:", isPrime)
}