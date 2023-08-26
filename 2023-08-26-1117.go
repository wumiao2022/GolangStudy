package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前时间
	fmt.Println(time.Now())

	// 练习2：计算矩形面积
	width := 5.0
	height := 3.0
	area := width * height
	fmt.Println("矩形的面积为：", area)

	// 练习3：字符串拼接
	str1 := "Hello"
	str2 := "World"
	str3 := str1 + " " + str2
	fmt.Println(str3)

	// 练习4：求和函数
	num1 := 10
	num2 := 20
	sum := add(num1, num2)
	fmt.Println("两个数的和为：", sum)
}

func add(a, b int) int {
	return a + b
}