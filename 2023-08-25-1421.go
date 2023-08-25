package main

import "fmt"

func main() {
	// 练习1: 输出Hello World
	fmt.Println("Hello World!")

	// 练习2: 定义变量并输出
	var name string = "Alice"
	fmt.Println("My name is", name)

	// 练习3: 数组遍历
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习4: 函数调用
	result := add(2, 3)
	fmt.Println("2 + 3 =", result)
}

func add(a, b int) int {
	return a + b
}