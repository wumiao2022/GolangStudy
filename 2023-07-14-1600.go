package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印"Hello, Golang!"
	fmt.Println("Hello, Golang!")

	// 练习2: 定义一个整型变量x并赋值为10，打印其值
	x := 10
	fmt.Println(x)

	// 练习3: 定义一个字符串变量name并赋值为"John"，打印其值
	name := "John"
	fmt.Println(name)

	// 练习4: 定义一个切片变量numbers并初始化包含1到5的整数，打印其值
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 练习5: 使用for循环打印1到10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6: 定义一个函数add，接收两个整型参数并返回它们的和，调用函数并打印结果
	fmt.Println(add(3, 5))

	// 练习7: 定义一个结构体Person，包含name和age字段，创建一个新的Person对象并打印其字段值
	person := Person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}