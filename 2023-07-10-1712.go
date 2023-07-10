package main

import "fmt"

func main() {
	// 实例1: 打印Hello World
	fmt.Println("Hello World")

	// 实例2: 定义并打印变量
	var num int
	num = 10
	fmt.Println("The value of num is:", num)

	// 实例3: 使用if条件语句
	if num > 5 {
		fmt.Println("num is greater than 5")
	} else {
		fmt.Println("num is less than or equal to 5")
	}

	// 实例4: 使用for循环打印数字
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 实例5: 使用函数计算两个数的和
	result := sum(5, 7)
	fmt.Println("The sum of 5 and 7 is:", result)
}

// 函数: 计算两个数的和
func sum(a, b int) int {
	return a + b
}