package main

import (
	"fmt"
	"time"
)

// 练习1：计算两个数的和
func addNumbers(a, b int) int {
	return a + b
}

func main() {
	// 练习2：输出当前日期和时间
	currentTime := time.Now()
	fmt.Println("Current date and time:", currentTime)

	// 练习3：判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}