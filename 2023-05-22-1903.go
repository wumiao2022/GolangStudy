package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1到100内的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2：打印出一个九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3：求出1到100内的所有质数
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}