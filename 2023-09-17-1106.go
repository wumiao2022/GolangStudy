package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3: 计算两个数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 练习4: 使用if语句判断一个数的正负性，并打印结果
	num := -5
	if num > 0 {
		fmt.Println("这是一个正数")
	} else if num < 0 {
		fmt.Println("这是一个负数")
	} else {
		fmt.Println("这是零")
	}

	// 练习5: 使用switch语句根据学生的分数判断其等级，并打印结果
	score := 80
	switch {
	case score >= 90:
		fmt.Println("等级为A")
	case score >= 80:
		fmt.Println("等级为B")
	case score >= 70:
		fmt.Println("等级为C")
	case score >= 60:
		fmt.Println("等级为D")
	default:
		fmt.Println("等级为E")
	}
}