package main

import (
	"fmt"
)

func main() {
	// 1. 输出10以内的偶数
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 2. 输出斐波那契数列前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}

	// 3. 打印乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 4. 计算1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 5. 判断一个数是不是质数
	n := 7
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(n, "是质数")
	} else {
		fmt.Println(n, "不是质数")
	}
}