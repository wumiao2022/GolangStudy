package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello, World!")

	// 练习2：定义并打印变量
	var number int = 10
	fmt.Println("The number is:", number)

	// 练习3：实现加法函数
	result := add(5, 3)
	fmt.Println("The sum is:", result)

	// 练习4：实现循环打印字符串
	str := "Golang"
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}

func add(a, b int) int {
	return a + b
}