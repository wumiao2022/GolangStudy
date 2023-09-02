package main

import "fmt"

func main() {
	// 练习1：将两个整数相加并打印结果
	a := 5
	b := 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习2：计算1到10的累加和并打印结果
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println("Total:", total)

	// 练习3：使用循环打印1到10的乘法表
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}

	// 练习4：判断一个数是否为素数
	num := 37
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}
}

请运行上述代码并观察结果。每日多练，加油！