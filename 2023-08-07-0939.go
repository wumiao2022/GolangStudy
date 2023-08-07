package main

import "fmt"

func main() {
	// 练习1: 打印出1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 打印出0到100之间的所有偶数
	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习3: 计算1到100之间的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4: 判断一个数是否是质数
	num := 29
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习5: 打印出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}