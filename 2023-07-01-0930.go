package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：将两个整数相加并打印结果
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：编写一个函数，接收两个参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 4)
	fmt.Println("Addition:", result)

	// 练习5：定义一个结构体类型Person，包含姓名和年龄字段
	type Person struct {
		Name string
		Age  int
	}

	// 创建一个Person对象并打印其字段值
	person := Person{Name: "Tom", Age: 25}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}