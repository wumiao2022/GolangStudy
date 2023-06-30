package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 打印出 1 到 10 的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// 练习2: 打印出 10 到 1 的整数
	for j := 10; j >= 1; j-- {
		fmt.Println(j)
	}

	fmt.Println()

	// 练习3: 打印出 1 到 10 的奇数
	for k := 1; k <= 10; k += 2 {
		fmt.Println(k)
	}

	fmt.Println()

	// 练习4: 打印出 1 到 10 的偶数
	for m := 2; m <= 10; m += 2 {
		fmt.Println(m)
	}

	fmt.Println()

	// 练习5: 打印出当前日期和时间
	fmt.Println("当前时间:", time.Now().Format("2006-01-02 15:04:05"))
}