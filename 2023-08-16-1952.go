package main

import "fmt"

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量x，并赋值为10，打印x的值
	x := 10
	fmt.Println(x)

	// 练习3：声明一个字符串变量name，并赋值为"John"，打印name的值
	name := "John"
	fmt.Println(name)

	// 练习4：定义一个切片s，包含整数1、2、3，打印s的长度和容量
	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s))

	// 练习5：定义一个函数add，接受两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))
}