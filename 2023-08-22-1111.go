package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2：计算1到10的累加和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("累加和:", sum)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习4：生成斐波那契数列
	num1, num2 := 0, 1
	fmt.Println("斐波那契数列：")
	for i := 0; i < 10; i++ {
		fmt.Println(num2)
		temp := num1 + num2
		num1 = num2
		num2 = temp
	}
}