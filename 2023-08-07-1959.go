package main

import "fmt"

func main() {
	// 练习一：打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习二：计算 1 到 100 的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习三：输出 1 到 100 中的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习四：输出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习五：计算斐波那契数列的前 10 个数字
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}