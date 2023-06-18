package main

import (
	"fmt"
)

func main() {
	// 1. 变量练习
	// 声明一个名为age的int型变量并赋初值18
	age := 18

	// 声明一个名为balance的float64型变量并不赋初值
	var balance float64

	// 声明一个名为name的string型变量并赋初值"Peter"
	name := "Peter"

	// 输出变量的值
	fmt.Println("年龄：", age)
	fmt.Println("余额：", balance)
	fmt.Println("姓名：", name)

	// 2. 数组练习
	// 创建一个长度为5的int数组，并赋初值
	nums := [5]int{1, 2, 3, 4, 5}

	// 输出数组的值
	fmt.Println("数组：", nums)

	// 3. 循环练习
	// for循环练习：输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// while循环练习：在数字小于5时循环输出Hello，world!
	j := 0
	for j < 5 {
		fmt.Println("Hello, world!")
		j++
	}

	// 4. 函数练习
	// 定义一个函数，接收两个int类型参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 调用函数并输出结果
	fmt.Println("1 + 2 = ", add(1, 2))
}