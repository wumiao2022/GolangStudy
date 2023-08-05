package main

import "fmt"

func main() {
	// 练习1：打印 Hello World
	fmt.Println("Hello, World!")

	// 练习2：计算并打印 1+2 的结果
	result := 1 + 2
	fmt.Println(result)

	// 练习3：定义一个变量，赋值为 "Golang"，然后打印出来
	language := "Golang"
	fmt.Println(language)

	// 练习4：使用 for 循环打印出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，接收两个参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	result = add(5, 3)
	fmt.Println(result)
}
