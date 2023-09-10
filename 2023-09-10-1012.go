package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello, World!" 到控制台
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量，赋值为10，并打印到控制台
	num := 10
	fmt.Println(num)

	// 3. 声明一个字符串变量，赋值为 "Go is awesome!"，并打印到控制台
	str := "Go is awesome!"
	fmt.Println(str)

	// 4. 声明一个布尔变量，赋值为 true，并打印到控制台
	isTrue := true
	fmt.Println(isTrue)

	// 5. 声明一个浮点数变量，赋值为 3.14，并打印到控制台
	floatNum := 3.14
	fmt.Println(floatNum)

	// 6. 声明一个整数变量 x，赋值为 5，和另一个整数变量 y，赋值为 10，将它们相加并打印到控制台
	x := 5
	y := 10
	sum := x + y
	fmt.Println(sum)

	// 7. 将上述代码包装在一个函数中，并调用该函数
	printExamples()
}

func printExamples() {
	fmt.Println("Hello, World!")
	fmt.Println(10)
	fmt.Println("Go is awesome!")
	fmt.Println(true)
	fmt.Println(3.14)

	x := 5
	y := 10
	sum := x + y
	fmt.Println(sum)
}