package main

import "fmt"

func main() {
	// 练习1 - 打印出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2 - 定义一个变量，存储你的名字，并打印出来
	name := "John"
	fmt.Println(name)

	// 练习3 - 定义两个整数变量，计算它们的和，并打印出来
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 练习4 - 使用循环打印出1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习5 - 创建一个切片，包含5个整数，并打印出来
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
}