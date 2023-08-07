package main

import (
	"fmt"
)

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算并打印两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：判断一个数是否为偶数
	num := 5
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：使用数组存储一组整数，并打印出数组中的每个元素
	arr := []int{1, 2, 3, 4, 5}
	for _, num := range arr {
		fmt.Println(num)
	}

	// 练习6：定义一个函数，输入两个数，返回它们的乘积
	num1 := 10
	num2 := 5
	product := multiply(num1, num2)
	fmt.Println("Product:", product)
}

func multiply(a, b int) int {
	return a * b
}