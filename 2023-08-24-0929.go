package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2：判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习3：打印Hello, World! 10次
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World!")
	}

	// 练习4：判断一个年份是否为闰年
	year := 2020
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("是闰年")
	} else {
		fmt.Println("不是闰年")
	}

	// 练习5：打印1到100之间的所有素数
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}