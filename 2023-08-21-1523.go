package main

import (
	"fmt"
)

func main() {
	// 练习1：打印出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：定义一个整型变量并赋值为10，打印出变量的值
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量并赋值为"Go语言"，打印出变量的值
	str := "Go语言"
	fmt.Println(str)

	// 练习4：定义一个浮点型变量并赋值为3.14, 打印出变量的值
	flt := 3.14
	fmt.Println(flt)

	// 练习5：定义一个字符串切片，包含元素"apple", "banana", "orange"，打印出切片的长度
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(len(fruits))
}