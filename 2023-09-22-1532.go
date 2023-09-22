package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：定义一个整数变量并输出
	num := 20
	fmt.Println(num)

	// 练习3：定义一个字符串变量并输出
	str := "Golang"
	fmt.Println(str)

	// 练习4：定义一个切片并输出
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习5：定义一个结构体并输出
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "John", Age: 25}
	fmt.Println(p)
}