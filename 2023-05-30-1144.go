package main

import (
	"fmt"
)

func main() {
	// 例子1：Hello World
	fmt.Println("Hello World")

	// 例子2：变量和常量
	var name string = "Golang"
	fmt.Println(name)
	const age int = 20
	fmt.Println(age)

	// 例子3：算术运算符
	var a int = 10
	var b int = 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// 例子4：条件语句
	var score int = 90
	if score >= 60 {
		fmt.Println("及格了")
	} else {
		fmt.Println("不及格")
	}

	// 例子5：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 例子6：函数
	var num1 int = 10
	var num2 int = 20
	fmt.Println(add(num1, num2))
}

func add(x, y int) int {
	return x + y
}