package main

import "fmt"

func main() {
	// 练习1：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习2：求出10的阶乘
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Println("10的阶乘是：", factorial)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习4：计算前n个斐波那契数
	n := 10
	a, b := 0, 1
	fmt.Print("前", n, "个斐波那契数：")
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}