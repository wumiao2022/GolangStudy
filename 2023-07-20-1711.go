package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量并赋值为10，打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：声明两个整数变量并分别赋值为5和7，打印两个变量的和
	var a int = 5
	var b int = 7
	sum := a + b
	fmt.Println(sum)

	// 练习4：声明一个字符串变量并赋值为"Go语言"，打印该变量的值
	str := "Go语言"
	fmt.Println(str)

	// 练习5：声明一个布尔变量并赋值为true，打印该变量的值
	isTrue := true
	fmt.Println(isTrue)
}