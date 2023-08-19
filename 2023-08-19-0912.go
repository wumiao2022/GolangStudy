package main

import "fmt"

func main() {
	// 练习1：将两个整数相加并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 练习2：将一个字符串反转并打印结果
	str := "Hello, Golang!"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)

	// 练习3：计算1到10之间的所有奇数之和并打印结果
	sumOdd := 0
	for i := 1; i <= 10; i += 2 {
		sumOdd += i
	}
	fmt.Println(sumOdd)

	// 练习4：判断一个数是否为素数并打印结果
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)
}