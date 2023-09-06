package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个变量，赋值为整数10，打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 3. 声明一个变量，赋值为浮点数3.14，打印该变量的值
	var pi float64 = 3.14
	fmt.Println(pi)

	// 4. 声明一个字符串变量，赋值为 "Golang"，打印该变量的值
	var str string = "Golang"
	fmt.Println(str)

	// 5. 声明一个布尔变量，赋值为true，打印该变量的值
	var isTrue bool = true
	fmt.Println(isTrue)

	// 6. 创建一个包含5个整数的数组，并打印数组的值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 7. 创建一个 map，包含姓名和年龄的键值对，并打印map的值
	person := map[string]int{
		"John": 25,
		"Jane": 30,
		"Bob":  40,
	}
	fmt.Println(person)

	// 8. 声明一个函数add，接受两个整数参数并返回它们的和，调用该函数并打印结果
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}
