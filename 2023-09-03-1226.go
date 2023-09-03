package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 输出当前日期和时间
	fmt.Println(time.Now())

	// 练习2: 输出从1到10的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习3: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}
