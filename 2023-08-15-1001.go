package main

import (
	"fmt"
)

func main() {
	// 1. 打印出 "Hello, World!" 的代码

	fmt.Println("Hello, World!")

	// ----------------------------------------------

	// 2. 计算两个整数的和并打印结果的代码

	num1 := 10
	num2 := 20
	sum := num1 + num2

	fmt.Println("Sum:", sum)

	// ----------------------------------------------

	// 3. 使用 for 循环打印出 1 到 10 的整数的代码

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// ----------------------------------------------

	// 4. 判断一个数是否为偶数并打印结果的代码

	num := 15

	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// ----------------------------------------------

	// 5. 创建一个切片（slice），添加一些元素并打印出来的代码

	slice := []int{1, 2, 3, 4, 5}

	for _, num := range slice {
		fmt.Println(num)
	}

	// ----------------------------------------------
}