package main

import (
	"fmt"
)

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算1-100之间的偶数和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 3. 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 4. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 5. 反转字符串
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)
}