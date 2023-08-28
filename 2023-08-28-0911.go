package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印所有小于100的奇数
	for i := 1; i < 100; i += 2 {
		fmt.Println(i)
	}

	// 练习2: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3: 判断一个数是不是素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 练习4: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}