package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 "Hello, World!" 到控制台
	fmt.Println("Hello, World!")

	// 练习2：声明一个整型变量 x，并初始化为 10，打印该变量的值到控制台
	var x int = 10
	fmt.Println(x)

	// 练习3：声明一个字符串变量 y，并初始化为 "Golang"，打印该变量的值到控制台
	var y string = "Golang"
	fmt.Println(y)

	// 练习4：声明一个切片变量 z，并初始化为包含 1, 2, 3 三个元素的切片，打印该变量的值到控制台
	z := []int{1, 2, 3}
	fmt.Println(z)

	// 练习5：定义一个函数 square，接收一个整型参数 x，并返回 x 的平方
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(5))

	// 练习6：使用 for 循环打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习7：使用 if-else 判断一个数是奇数还是偶数，并打印结果到控制台
	num := 9
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}