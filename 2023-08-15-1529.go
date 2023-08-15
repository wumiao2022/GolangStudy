package main

import "fmt"

func main() {
	// 1. 计算两个数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 2. 判断一个数是否为偶数
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 3. 打印1到10的平方数
	for i := 1; i <= 10; i++ {
		square := i * i
		fmt.Println(square)
	}

	// 4. 定义一个函数，输入两个数，返回较大的数
	max := findMax(15, 20)
	fmt.Println("Max:", max)
}

func findMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}