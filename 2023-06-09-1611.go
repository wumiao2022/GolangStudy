package main

import "fmt"

func main() {
	// 练习一：输出 Hello World
	fmt.Println("Hello World")

	// 练习二：计算并输出 1 + 2 的结果
	fmt.Println(1 + 2)

	// 练习三：定义并输出一个整型变量x，赋值为10
	x := 10
	fmt.Println(x)

	// 练习四：定义一个名为add的函数，接受两个参数x和y，返回它们的和
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}