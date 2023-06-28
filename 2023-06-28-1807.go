package main

import "fmt"

func main() {
	// 练习1：输出 Hello World
	fmt.Println("Hello World")

	// 练习2：输出数字相加结果
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println(num1 + num2)

	// 练习3：循环输出数字 1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：判断两个数的大小关系
	var a, b int
	a = 5
	b = 8
	if a > b {
		fmt.Println("a 大于 b")
	} else if a < b {
		fmt.Println("a 小于 b")
	} else {
		fmt.Println("a 等于 b")
	}

	// 练习5：使用 struct 定义一个学生信息，并输出
	type student struct {
		name string
		age  int
	}
	s := student{name: "John", age: 18}
	fmt.Println(s)
}