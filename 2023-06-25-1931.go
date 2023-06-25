package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出1~100之间的偶数
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习2: 计算1~100之间所有奇数的和
	sum := 0
	for i := 1; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println("1~100之间所有奇数的和为：", sum)

	// 练习3: 判断一个数是否为素数
	num := 17
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

	// 练习4: 逆序输出一个字符串
	str := "hello world"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
}