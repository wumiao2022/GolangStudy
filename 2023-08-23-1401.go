package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 "Hello, world!"
	fmt.Println("Hello, world!")

	// 练习2：打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1到10的所有整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：定义一个函数，接受两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Sum of 3 and 4:", add(3, 4))

	// 练习5：定义一个结构体Person，包含姓名和年龄字段，并创建一个Person类型的变量
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(person)
}