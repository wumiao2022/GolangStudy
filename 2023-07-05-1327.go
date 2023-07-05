package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数，并输出结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4: 遍历打印1到100之间的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习5: 定义一个函数，实现两数相乘的功能，并输出结果
	multiply(5, 6)
}

func multiply(x, y int) {
	result := x * y
	fmt.Println("Result:", result)
}