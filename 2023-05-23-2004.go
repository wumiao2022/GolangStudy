package main

import (
	"fmt"
)

func main() {
	// 1. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 2. 判断一个年份是否为闰年
	year := 2020
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%d是闰年", year)
	} else {
		fmt.Printf("%d不是闰年", year)
	}

	// 3. 求斐波那契数列的第n项
	n := 10
	fmt.Printf("斐波那契数列的第%d项是%d", n, fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}