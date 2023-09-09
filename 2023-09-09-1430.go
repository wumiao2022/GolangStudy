package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：输出数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个数是奇数还是偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：计算一个数的平方
	number := 6
	square := number * number
	fmt.Println("平方:", square)

	// 练习5：计算两个数的和
	num1 := 3
	num2 := 4
	sum := num1 + num2
	fmt.Println("和:", sum)
}