package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并给其赋值为10，然后打印这个变量的值
	var num int = 10
	fmt.Println(num)

	// 3. 声明两个整数变量，分别为x和y，并交换它们的值，然后打印交换后的结果
	x := 5
	y := 8
	x, y = y, x
	fmt.Println("x =", x, "y =", y)

	// 4. 声明一个字符串变量，获取用户输入的值，并打印出来
	var input string
	fmt.Print("请输入一个字符串：")
	fmt.Scanln(&input)
	fmt.Println("你输入的字符串是：", input)

	// 5. 声明一个数组并初始化，包含5个整数元素，然后打印这个数组的长度
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组的长度为：", len(arr))
}