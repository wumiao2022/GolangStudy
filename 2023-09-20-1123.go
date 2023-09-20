package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和并打印结果
	a := 3
	b := 5
	sum := a + b
	fmt.Println(sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 13
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，接收两个整数参数并返回它们的乘积
	multiply(2, 4)
}

func multiply(a, b int) {
	fmt.Println(a * b)
}