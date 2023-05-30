package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并输出12*24的结果
	result := 12 * 24
	fmt.Println(result)

	// 练习3：创建一个整数变量并给其赋值，再将其输出
	num := 5
	fmt.Println(num)

	// 练习4：创建一个字符串变量并给其赋值，再将其输出
	str := "Hello, Golang!"
	fmt.Println(str)

	// 练习5：使用循环输出数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6：创建一个函数，接收两个整数并返回它们的和
	fmt.Println(add(5, 3))

	// 练习7：创建一个切片并将其中的元素输出
	slice := []string{"apple", "banana", "orange"}
	fmt.Println(slice)
}

func add(a, b int) int {
	return a + b
}