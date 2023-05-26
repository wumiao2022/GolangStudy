package main

import "fmt"

func main() {
	// 练习1：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Printf("\n")
	}

	// 练习2：计算斐波那契数列的前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Printf("\n")

	// 练习3：判断一个数是否是素数
	num := 23
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}

	// 练习4：求一个整数的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Printf("%d的阶乘是%d\n", n, factorial)
}