package main

import (
	"fmt"
)

func main() {
	// 1. 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算整数相加结果，并打印
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 判断一个整数是否是偶数，并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 4. 使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 5. 使用for循环打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
