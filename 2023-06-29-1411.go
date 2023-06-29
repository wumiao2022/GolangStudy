package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：定义一个整数变量并打印
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量并打印
	str := "Golang"
	fmt.Println(str)

	// 练习4：定义一个浮点数变量并打印
	floatNum := 3.14
	fmt.Println(floatNum)

	// 练习5：定义一个布尔变量并打印
	flag := true
	fmt.Println(flag)

	// 练习6：定义一个数组并打印
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习7：定义一个切片并打印
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习8：定义一个Map并打印
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)

	// 练习9：定义一个函数并调用
	printHello()
}

func printHello() {
	fmt.Println("Hello")
}
