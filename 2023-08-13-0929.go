package main

import "fmt"

func main() {
	// 练习1：打印出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：打印出1到100之间所有的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习3：计算1到100之间所有奇数的总和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("奇数总和：", sum)

	// 练习4：输出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
