package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// 练习2：打印10-1的数字
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// 练习3：打印1-10中的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()

	// 练习4：计算1-10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("1-10的和为：%d", sum)

	fmt.Println()

	// 练习5：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数", num)
	} else {
		fmt.Printf("%d不是素数", num)
	}
}