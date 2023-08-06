package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义并打印一个字符串变量
	message := "Hello, Golang!"
	fmt.Println(message)

	// 3. 定义并打印一个整数变量
	number := 42
	fmt.Println(number)

	// 4. 定义一个函数，接收两个整数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 5. 调用上面定义的函数，并打印结果
	result := add(3, 5)
	fmt.Println(result)
}