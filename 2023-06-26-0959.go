package main

import "fmt"

func main() {
	// 练习1：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和为：", sum)

	// 练习2：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习3：判断一个数是否为素数
	num := 13
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
}