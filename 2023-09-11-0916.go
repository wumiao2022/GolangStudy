package main

import (
	"fmt"
)

func main() {
	// 练习1：将字符串反转
	str := "hello world"
	reverseStr := ""
	for _, s := range str {
		reverseStr = string(s) + reverseStr
	}
	fmt.Println(reverseStr)

	// 练习2：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习3：计算Fibonacci数列的第n个数
	n := 10
	fibonacci := make([]int, n+1)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println(fibonacci[n])
}