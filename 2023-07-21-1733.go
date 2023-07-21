package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个字符串变量并打印
	name := "John"
	fmt.Println(name)

	// 练习3: 定义一个整数变量并打印
	age := 25
	fmt.Println(age)

	// 练习4: 使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 定义一个函数，接收两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	// 练习6: 定义一个结构体并初始化
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(person)
}