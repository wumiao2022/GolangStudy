package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1+2的结果
	fmt.Println(1 + 2)

	// 练习3：将一个整数变量的值加1，并打印结果
	num := 10
	num++
	fmt.Println(num)

	// 练习4：使用for循环打印1到5的数字
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习5：声明一个字符串变量，赋予值"Go语言编程"，并打印
	str := "Go语言编程"
	fmt.Println(str)

	// 练习6：定义一个函数，接收两个整数参数，返回它们的和
	sum := add(2, 3)
	fmt.Println(sum)
}

// 定义一个函数add，接收两个整数参数，返回它们的和
func add(x, y int) int {
	return x + y
}