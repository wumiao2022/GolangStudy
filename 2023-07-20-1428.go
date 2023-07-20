package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Printf("当前时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// 打印"Hello, Golang!" 5次
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, Golang!")
	}

	// 计算并打印1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("1到100之间所有偶数的和为：%d\n", sum)

	// 打印1到10之间的乘法表
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}