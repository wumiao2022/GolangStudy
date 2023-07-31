package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明和初始化一个整型变量x，并将其赋值为10
	x := 10
	fmt.Println(x)

	// 3. 声明一个字符串变量name，并将其赋值为"John"
	name := "John"
	fmt.Println(name)

	// 4. 定义一个函数add，接收两个整型参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	sum := add(3, 4)
	fmt.Println(sum)

	// 5. 声明一个整型常量PI，并将其赋值为3.14
	const PI = 3.14
	fmt.Println(PI)
}
