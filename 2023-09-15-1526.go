package main

import (
	"fmt"
)

func main() {
	//练习1：打印Hello World
	fmt.Println("Hello World")

	//练习2：计算两个整数的和并打印结果
	a, b := 10, 20
	sum := a + b
	fmt.Println(sum)

	//练习3：计算两个浮点数的乘积并打印结果
	x, y := 3.14, 2.5
	prod := x * y
	fmt.Println(prod)

	//练习4：定义一个字符串变量并打印其长度
	str := "Hello, Golang"
	fmt.Println(len(str))

	//练习5：定义一个切片并打印其元素个数和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice), cap(slice))

	//练习6：定义一个结构体来表示一个人的信息，并打印该结构体变量的值
	type person struct {
		name string
		age  int
	}
	p := person{name: "Alice", age: 25}
	fmt.Println(p)
}