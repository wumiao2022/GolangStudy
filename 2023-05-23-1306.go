package main

import (
"fmt"
)

func main() {
	// 练习1：打印出1~100内的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // 换行

	// 练习2：for循环嵌套打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println() // 换行
	}

	// 练习3：从键盘输入一个数n，计算1到n的和
	var n int
	var sum int
	fmt.Printf("请输入一个正整数n：")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1到%d的和是%d\n", n, sum)
}