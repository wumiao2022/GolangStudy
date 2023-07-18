package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数相加的结果并打印
	num1 := 10
	num2 := 20
	result := num1 + num2
	fmt.Println(result)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 3
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}