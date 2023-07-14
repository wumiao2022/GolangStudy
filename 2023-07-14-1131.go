package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算并打印1+2的结果
	fmt.Println(1 + 2)

	// 3. 定义一个整数变量x，赋值为10，并打印出来
	x := 10
	fmt.Println(x)

	// 4. 定义一个字符串变量name，赋值为"John"，并打印出来
	name := "John"
	fmt.Println(name)

	// 5. 定义一个函数double，接收一个整数作为参数，返回该整数乘以2的结果
	double := func(n int) int {
		return n * 2
	}

	// 6. 调用double函数，传入参数5，并打印出结果
	fmt.Println(double(5))
}