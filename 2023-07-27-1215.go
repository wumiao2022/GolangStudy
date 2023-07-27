package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整数变量，赋值为10，并打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，赋值为"Go语言"，并打印该变量的值
	var str string = "Go语言"
	fmt.Println(str)

	// 练习4：定义一个切片，初始化包含1，2，3三个元素，并打印该切片的长度和容量
	slice := []int{1, 2, 3}
	fmt.Println("长度:", len(slice), "容量:", cap(slice))

	// 练习5：定义一个结构体类型Person，包含姓名和年龄字段，创建一个Person类型的变量，并打印其值
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 20}
	fmt.Println(p)
}