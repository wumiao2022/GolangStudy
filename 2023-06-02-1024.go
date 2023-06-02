package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello, world!" 到控制台
	fmt.Println("Hello, world!")

	// 2. 定义一个名为 x 的整数变量，赋值为 5
	x := 5
	fmt.Println(x)

	// 3. 定义一个名为 s 的字符串变量，赋值为 "Golang"
	s := "Golang"
	fmt.Println(s)

	// 4. 定义一个名为 nums 的整数切片，包含 1、2、3、4、5
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 5. 定义一个名为 person 的结构体类型，包含 name 和 age 两个字段
	type person struct {
		name string
		age  int
	}
	p := person{name: "Alice", age: 30}
	fmt.Println(p)
}