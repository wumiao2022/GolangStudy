package main

import "fmt"

func main() {
	// 1. 反转字符串
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 2. 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 3. 求斐波那契数列的第n个数
	n := 10
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println(fibonacci[n-1])
}