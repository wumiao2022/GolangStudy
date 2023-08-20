package main

import (
	"fmt"
)

func main() {
	// 实例1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 实例3：计算一个整数的平方
	num := 5
	square := num * num
	fmt.Println("Square:", square)

	// 实例4：判断一个整数是否为偶数
	num2 := 6
	if num2%2 == 0 {
		fmt.Println(num2, "is even.")
	} else {
		fmt.Println(num2, "is odd.")
	}

	// 实例5：打印1到10之间（包括10）的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 实例6：计算阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, "is", factorial)
}