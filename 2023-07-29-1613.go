package main

import "fmt"

func main() {
	// 练习1：打印"Hello, Go!"
	fmt.Println("Hello, Go!")

	// 练习2：计算并打印 5 加 3 的结果
	fmt.Println(5 + 3)

	// 练习3：声明一个字符串变量，赋值为 "Golang"，并打印出来
	str := "Golang"
	fmt.Println(str)

	// 练习4：声明一个整数变量，赋值为 10，再声明一个浮点数变量，赋值为 3.5，将两个变量相乘并打印结果
	num1 := 10
	num2 := 3.5
	fmt.Println(num1 * int(num2))

	// 练习5：声明一个布尔变量，赋值为 true，再声明一个整数变量，赋值为 8，如果布尔变量为 true，则将整数加 5，否则将整数减 3，并打印结果
	flag := true
	num3 := 8
	if flag {
		fmt.Println(num3 + 5)
	} else {
		fmt.Println(num3 - 3)
	}
}