package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello, World!")

	// 练习2：计算并输出1+2的结果
	fmt.Println(1 + 2)

	// 练习3：声明并初始化一个整型变量，输出其值
	num := 10
	fmt.Println(num)

	// 练习4：声明一个字符串变量并给其赋值"Go is great!"
	str := "Go is great!"
	fmt.Println(str)

	// 练习5：利用for循环输出1到5的所有数字
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习6：声明一个切片，并循环遍历输出每个元素的值
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}

	// 练习7：实现一个计算两个数相加的函数，并调用该函数输出结果
	sum := add(3, 4)
	fmt.Println(sum)
}

// 练习7所需的函数
func add(a, b int) int {
	return a + b
}