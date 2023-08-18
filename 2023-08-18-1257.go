package main

import (
	"fmt"
)

func main() {
	// 示例1：输出Hello World
	fmt.Println("Hello World")

	// 示例2：变量赋值和打印
	var num1 int = 10
	var num2 int = 20
	var sum int = num1 + num2
	fmt.Println("Sum:", sum)

	// 示例3：循环打印数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例4：切片和循环打印
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 示例5：条件语句和判断
	score := 90
	if score >= 80 {
		fmt.Println("Excellent!")
	} else if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// 示例6：函数和返回值
	result := multiply(5, 2)
	fmt.Println("Result:", result)
}

func multiply(a, b int) int {
	return a * b
}