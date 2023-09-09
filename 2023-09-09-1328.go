package main

import (
	"fmt"
)

func main() {
	// 1. 输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并赋值为10，然后输出该变量的值
	num := 10
	fmt.Println(num)

	// 3. 声明一个字符串变量并赋值为"GoLang"，然后输出该变量的值
	str := "GoLang"
	fmt.Println(str)

	// 4. 声明一个浮点数变量并赋值为3.14，然后输出该变量的值
	fl := 3.14
	fmt.Println(fl)

	// 5. 声明一个布尔型变量并赋值为true，然后输出该变量的值
	b := true
	fmt.Println(b)

	// 6. 声明一个整型数组，元素个数为5，然后输出数组的长度
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
}
