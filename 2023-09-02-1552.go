package main

import "fmt"

func main() {
	// 练习1：打印 Hello World
	fmt.Println("Hello World")

	// 练习2：循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个整数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习4：定义一个函数，接收一个整数参数，判断是否为偶数并返回结果
	isEven := func(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	}
	fmt.Println(isEven(7)) // false
	fmt.Println(isEven(10)) // true

	// 练习5：定义一个结构体类型，包含姓名和年龄字段，并创建一个结构体实例并打印其内容
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"John Doe", 25}
	fmt.Println(p.Name, p.Age)
}