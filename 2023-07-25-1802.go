package main

import "fmt"

func main() {
	// 1. 展示如何定义一个整数类型变量，并将其赋值为任意整数值
	var num int
	num = 10

	// 2. 展示如何定义一个浮点数类型变量，并将其赋值为任意浮点数值
	var fNum float64
	fNum = 3.14

	// 3. 展示如何定义一个布尔类型变量，并将其赋值为 true 或 false
	var isTrue bool
	isTrue = true

	// 4. 展示如何定义一个字符串类型变量，并将其赋值为任意字符串值
	var str string
	str = "Hello, Go!"

	// 5. 展示如何定义一个切片类型变量，并将其初始化为包含 3 个整数元素的切片
	var slice []int
	slice = []int{1, 2, 3}

	// 6. 展示如何定义一个映射类型变量，并将其初始化为包含两个键值对的映射
	var m map[string]int
	m = map[string]int{"apple": 1, "banana": 2}

	// 7. 展示如何定义一个结构体类型变量，并将其初始化为具有相应字段值的结构体
	type person struct {
		name string
		age  int
	}
	var p person
	p = person{name: "Alice", age: 20}

	fmt.Println(num, fNum, isTrue, str, slice, m, p)
}