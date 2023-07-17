package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 声明一个整数变量并打印其值
	var num int = 10
	fmt.Println(num)

	// 练习3: 声明一个字符串变量并打印其值
	var str string = "Golang"
	fmt.Println(str)

	// 练习4: 计算两个整数的和并打印结果
	var a, b int = 5, 7
	sum := a + b
	fmt.Println(sum)

	// 练习5: 使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6: 使用if语句判断一个数字是否为偶数，并打印结果
	num := 8
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}